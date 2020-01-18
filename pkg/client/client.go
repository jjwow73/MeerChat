package client

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

type Client struct {
	Conn *websocket.Conn

	Send chan []byte // Message: server -> client
}

func DoChatting() {
	handleInput()
}

func connectToWebsocket(addr string, id string, password string) (conn *websocket.Conn, err error) {
	query := "id=" + id + "&" + "password=" + password
	u := url.URL{Scheme: "ws", Host: addr, Path: "/ws", RawQuery: query}
	log.Printf("connecting to %s", u.String())

	conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	return conn, err
}

func readMessage(room *room) {
	for {
		_, message, err := room.conn.ReadMessage()
		if err != nil {
			log.Println("room ", room.id, " read error:", err)
			close(room.connErr)
			return
		}
		if room.ifFocused() {
			log.Printf("recv: %s", message)
		}
	}
}

func writeMessage(message string) {
	room, exist := getRoom(*roomList.focusedId)
	if !exist {
		log.Println("non entered-room")
		return
	}
	err := room.conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("write error: ", err)
		close(room.done)
	}
}
