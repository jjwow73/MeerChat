package main

import (
	"../../pkg/client"
	"flag"
)

var (
	addr     = flag.String("addr", "localhost:8080", "http service address")
	id       = flag.String("id", "1", "A id of a room")
	password = flag.String("password", "1", "A password of a room")
)

func main() {
	client.RpcStart()
	//client.Start()
}
