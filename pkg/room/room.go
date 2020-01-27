package room

import (
	"../User"
)

type Room struct {
	name string
	admin User.User
	password string
	description string
	ipAddress string
	portNumber string
	maxUserNumber int32
	userList []User.User
	blackList []User.User
	updatedAt int32
}

func NewRoom(_name string, _description string, _password string, _ipAddress string, _portNumber string) *Room {

	return &Room{description: _description}
}