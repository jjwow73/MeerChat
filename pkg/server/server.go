package server

import (
	"../client"

	"github.com/gorilla/websocket"
	"log"
	"net/http"
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
	room, _ := getRoom(id, password)

	client := &client.Client{Conn: conn, Send: make(chan []byte, 256)}
	room.register(client)

	go sendMessageToClient(room, client)
	receiveMessageFromClient(room, client)
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

func sendMessageToClient(room *Room, client *client.Client) {
	defer room.unregister(client)
	for {
		message, ok := <-client.Send
		if !ok {
			log.Println("channel closed")
			client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		err := client.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			room.broadcast([]byte (err.Error()))
			log.Println("write error:", err)
			return
		}
	}
}

func receiveMessageFromClient(room *Room, client *client.Client) {
	defer room.unregister(client)
	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}
		log.Printf("message: %s", message)

		room.broadcast(message)
	}
}
