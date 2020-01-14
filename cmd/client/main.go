package main

import (
	"../../pkg/client"

	"flag"
	"log"
)

var (
	port = flag.String("port", "9000", "port used for ws connection")
)

func main() {
	flag.Parse()
	// connect
	ws, err := client.Client(port)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	// receive
	go client.ReceiveMessage(ws)
	// send
	client.SendMessage(ws)
}
