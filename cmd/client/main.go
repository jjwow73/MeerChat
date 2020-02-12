package main

import (
	"flag"

	"github.com/jjwow73/MeerChat/pkg/client"
)

var (
	addr     = flag.String("addr", "localhost:8080", "http service address")
	id       = flag.String("id", "1", "A id of a room")
	password = flag.String("password", "1", "A password of a room")
)

func main() {

	go client.CuiMain()
	client.RpcStart()
	//client.Start()
}
