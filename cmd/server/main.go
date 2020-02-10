package main

import (
	"flag"

	"../../pkg/server"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()
	server.Start(*addr)
}
