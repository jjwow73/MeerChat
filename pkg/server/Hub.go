package server

import (
	"../message"

	"fmt"
	"golang.org/x/net/websocket"
)

type hub struct {
	clients          map[string]*websocket.Conn
	addClientChan    chan *websocket.Conn
	removeClientChan chan *websocket.Conn
	broadcastChan    chan message.Message
}

func NewHub() *hub {
	return &hub{
		clients:          make(map[string]*websocket.Conn),
		addClientChan:    make(chan *websocket.Conn),
		removeClientChan: make(chan *websocket.Conn),
		broadcastChan:    make(chan message.Message),
	}
}

func (h *hub) run() {
	for {
		select {
		case conn := <-h.addClientChan:
			h.addClient(conn)
		case conn := <-h.removeClientChan:
			h.removeClient(conn)
		case m := <-h.broadcastChan:
			h.broadcastMessage(m)
		}
	}
}

func (h *hub) removeClient(conn *websocket.Conn) {
	delete(h.clients, conn.LocalAddr().String())
}
func (h *hub) addClient(conn *websocket.Conn) {
	h.clients[conn.RemoteAddr().String()] = conn
}

func (h *hub) broadcastMessage(m message.Message) {
	for _, conn := range h.clients {
		err := websocket.JSON.Send(conn, m)
		if err != nil {
			fmt.Println("Error broadcasting message: ", err)
			return
		}
	}
}
