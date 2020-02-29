package model

import "github.com/gorilla/websocket"

type Room struct {
	id string
	password string
	ip string
	port string
	conn *websocket.Conn
}

func (room *Room) createRoom() {

}
