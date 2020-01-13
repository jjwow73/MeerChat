package main

import (
	"../../pkg"

	"flag"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

var (
	port = flag.String("port", "9000", "port used for ws connection")
)

func server(port string) error {
	h := pkg.NewHub()
	mux := http.NewServeMux()
	mux.Handle("/", websocket.Handler(func(ws *websocket.Conn) {
		pkg.Handler(ws, h)
	}))
	s := http.Server{Addr: ":" + port, Handler: mux}
	return s.ListenAndServe()
}

func main() {
	flag.Parse()
	log.Fatal(server(*port))
}
