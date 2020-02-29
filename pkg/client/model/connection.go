package model

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

type connection struct {
	conn *websocket.Conn
}

func (c *connection) join(room *Room, user *User) (*websocket.Conn, error) {
	query := "id=" + room.id + "&password=" + room.password + "&name=" + user.name
	u := url.URL{Scheme: "ws", Host: room.ip + ":" + room.port, Path: "/ws", RawQuery: query}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	return conn, err
}
