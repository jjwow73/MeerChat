package main

import (
	"../../pkg/client"

	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var (
	addr     = flag.String("addr", "localhost:8080", "http service address")
	id       = flag.String("id", "1", "A id of a room")
	password = flag.String("password", "1", "A password of a room")
)

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	conn, err := client.ConnectToWebsocket(addr, id, password)
	if err != nil {
		log.Fatal("dial:", err)
		return
	}
	defer conn.Close()

	done := make(chan struct{})
	go client.ReadMessage(conn, done)
	go client.WriteMessage(conn, done)

	select {
	case <-done:
		log.Println("connection is broken...")
		return
	case <-interrupt:
		log.Println("interrupt")

		err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			log.Println("write close:", err)
			return
		}
		select {
		case <-done:
		case <-time.After(time.Second):
		}
		return
	}
}
