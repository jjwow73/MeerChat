package main

import (
	"../../pkg/server"

	"flag"
	"log"
)

var (
	port = flag.String("port", "9000", "port used for ws connection")
)

func main() {
	flag.Parse()
	log.Fatal(server.Server(*port))
}
