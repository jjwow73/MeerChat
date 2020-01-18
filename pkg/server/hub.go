package server

import (
	"github.com/gorilla/websocket"
	"log"
)

type Hub struct {
	clients    map[*clientInfo]bool
	broadcast  chan []byte
	register   chan *clientInfo
	unregister chan *clientInfo
	active     chan bool
}

type clientInfo struct {
	conn *websocket.Conn
	send chan []byte // Message: server -> client
}

func newHub() *Hub {
	hub := &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *clientInfo),
		unregister: make(chan *clientInfo),
		clients:    make(map[*clientInfo]bool),
		active:     make(chan bool),
	}
	return hub
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			log.Println("unregister occurred")
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				client = nil
				if len(h.clients) == 0 {
					log.Println("no client... remove hub")
					h.active <- false
				}
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
