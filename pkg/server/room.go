package server

import (
	"../chat"
	"github.com/gorilla/websocket"
	"log"
)

const unAuthMessage = "meer"

type room struct {
	id       string
	password string
	hub      *hub
}

var rooms = make(map[string]*room)

func createRoom(id string, password string) *room {
	hub := newHub()
	room := &room{id: id, password: password, hub: hub}
	rooms[id] = room

	// If room has no connection then remove room
	go func() {
		<-hub.done
		log.Println("no connection in room. delete room", room.id)
		delete(rooms, room.id)
		room = nil
	}()
	return room
}

func getRoom(id string, password string) (room *room, auth bool) {
	if room, exist := rooms[id]; exist {
		return room, password == room.password
	}
	return createRoom(id, password), true
}

func (room *room) broadcast(message *chat.Message) {
	room.hub.broadcast <- message
}

func (room *room) register(connInfo *connInfo) {
	room.hub.register <- connInfo
}

func (room *room) unregister(connInfo *connInfo) {
	room.hub.unregister <- connInfo
}

func (room *room) receiveMessage(connInfo *connInfo) {
	for {
		_, message, err := connInfo.conn.ReadMessage()
		if !connInfo.auth {
			message = []byte(unAuthMessage)
		}
		if err != nil {
			log.Println("read error:", err)
			return
		}
		log.Println("chat:", string(message))

		room.broadcast(&chat.Message{Content: message, Name: connInfo.clientName})
	}
}

func (room *room) sendMessage(connInfo *connInfo) {
	for {
		message, ok := <-connInfo.channel
		if !connInfo.auth {
			message.Content = []byte(unAuthMessage)
		}
		if !ok {
			log.Println("connection closed")
			connInfo.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}
		err := connInfo.conn.WriteJSON(&message)
		if err != nil {
			log.Println("write error:", err)
			return
		}
	}
}
