package server

import (
	"../chat"
	"github.com/gorilla/websocket"
)

type connInfo struct {
	conn       *websocket.Conn
	auth       bool
	clientName string
	channel    chan *chat.Message
}

func newConnInfo(conn *websocket.Conn, auth bool, name string) *connInfo {
	return &connInfo{
		conn:       conn,
		auth:       auth,
		clientName: name,
		channel:    make(chan *chat.Message, 256),
	}
}
