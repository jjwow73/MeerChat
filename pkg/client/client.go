package client

import (
	"bufio"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
)

type Client struct {
	Conn *websocket.Conn

	Send chan []byte	// Message: server -> client
}

func ConnectToWebsocket(addr *string, id *string, password *string) (conn *websocket.Conn, err error) {
	query := "id=" + *id + "&" + "password=" + *password
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws", RawQuery: query}
	log.Printf("connecting to %s", u.String())

	conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	return conn, err
}

func ReadMessage(conn *websocket.Conn, done chan struct{}) {
	defer close(done)
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			return
		}
		log.Printf("recv: %s", message)
	}
}

func WriteMessage(conn *websocket.Conn, done chan struct{}) {
	defer close(done)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		err := conn.WriteMessage(websocket.TextMessage, []byte(text))
		if err != nil {
			log.Println("write error:", err)
			break
		}
	}
}
