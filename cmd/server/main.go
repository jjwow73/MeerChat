package main

import (
	"../../pkg"

	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()
	hub := pkg.NewHub()
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		pkg.ServeHandler(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
