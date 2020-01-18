package server

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

const (
	meer = "meer"
	meerModeMessage = "You've got wrong password. Enter to Meerkat mode."
)

func ServeHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	id, password, ok := getParamsFromUrl(r)
	if !ok {
		return
	}
	room, auth := getRoom(id, password)
	if !auth {
		conn.WriteMessage(websocket.TextMessage, []byte(meerModeMessage))
	}

	client := &clientInfo{conn: conn, send: make(chan []byte, 256)}
	room.register(client)
	defer room.unregister(client)

	go sendMessageToClient(room, client, auth)
	receiveMessageFromClient(room, client, auth)
}

func getParamsFromUrl(r *http.Request) (id string, password string, ok bool) {
	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return id, password, ok
	}
	id = ids[0]
	passwords, ok := r.URL.Query()["password"]
	if !ok || len(ids[0]) < 1 {
		log.Println("Url Param 'password' is missing")
		return id, password, ok
	}
	password = passwords[0]
	return id, password, ok
}

func sendMessageToClient(room *Room, client *clientInfo, auth bool) {
	for {
		message, ok := <-client.send
		if !auth {
			message = []byte(meer)
		}
		if !ok {
			log.Println("channel closed")
			client.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		err := client.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			room.broadcast([]byte (err.Error()))
			log.Println("write error:", err)
			return
		}
	}
}

func receiveMessageFromClient(room *Room, client *clientInfo, auth bool) {
	for {
		_, message, err := client.conn.ReadMessage()
		if !auth {
			message = []byte(meer)
		}
		if err != nil {
			log.Println("read error:", err)
			break
		}
		log.Printf("message: %s", message)

		room.broadcast(message)
	}
}
