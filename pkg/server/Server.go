package server

import (
	"../message"

	"golang.org/x/net/websocket"
	"net/http"
)

func Server(port string) error {
	h := NewHub()
	mux := http.NewServeMux()
	mux.Handle("/", websocket.Handler(func(ws *websocket.Conn) {
		handler(ws, h)
	}))
	s := http.Server{Addr: ":" + port, Handler: mux}
	return s.ListenAndServe()
}

func handler(ws *websocket.Conn, h *hub) {
	go h.run()
	h.addClientChan <- ws
	for {
		var m message.Message
		err := websocket.JSON.Receive(ws, &m)
		if err != nil {
			h.broadcastChan <- message.Message{err.Error()}
			h.removeClient(ws)
			return
		}
		h.broadcastChan <- m
	}
}
