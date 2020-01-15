package server

import (
	"../client"

	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// serveWs handles websocket requests from the peer.
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
	room, _ := GetRoom(id, password)
	hub := room.Hub

	client := &client.Client{Conn: conn, Send: make(chan []byte, 256)}
	hub.register <- client

	go func() {
		for {
			message, ok := <-client.Send
			if !ok {
				// The hub closed the channel.
				conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err = conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				hub.broadcast <- []byte(err.Error())
				hub.unregister <- client
				log.Println("write:", err)
				return
			}
		}
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("message: %s", message)

		hub.broadcast <- message
	}
}

func getParamsFromUrl(r *http.Request) (id string, password string, ok bool){
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
