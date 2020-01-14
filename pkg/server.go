package pkg

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// serveWs handles websocket requests from the peer.
func ServeHandler(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{conn: conn, send: make(chan []byte, 256)}
	hub.register <- client

	go func() {
		for {
			message, ok := <-client.send
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

	//go client.WritePump()

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
