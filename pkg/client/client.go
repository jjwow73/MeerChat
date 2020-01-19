package client

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

type Client struct {
	Conn *websocket.Conn

	Send chan []byte // Message: server -> client
}

func DoChatting() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go handleInput()

	select {
	case <-interrupt:
		leaveAllRoom()
		<-time.After(time.Second)
	}
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
			select {
			case <-room.done:	// normal closed
			default:			// abnormal closed
				log.Println("room", room.id, " read error:", err)
				close(room.done)
			}
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
