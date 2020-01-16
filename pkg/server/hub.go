package server

import (
	"../client"
	"context"
	"log"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Hub struct {
	clients    map[*client.Client]bool
	broadcast  chan []byte
	register   chan *client.Client
	unregister chan *client.Client
	active     chan bool
}

func newHub() *Hub {
	hub := &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *client.Client),
		unregister: make(chan *client.Client),
		clients:    make(map[*client.Client]bool),
		active:     make(chan bool),
	}
	return hub
}

func (h *Hub) run(ctx context.Context) {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			log.Println("unregister occurred")
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
				client = nil
				if len(h.clients) == 0 {
					log.Println("no client... remove hub")
					h.active <- false
				}
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.clients, client)
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
