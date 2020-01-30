package room

import (
	"github.com/wkd3475/MeerChat/pkg/user"
)

type Room struct {
	name          string
	admin         user.User
	password      string
	description   string
	ipAddress     string
	portNumber    string
	maxUserNumber int32
	userList      []user.User
	blackList     []user.User
	updatedAt     int32
}

func NewRoom(_name string, _description string, _password string, _ipAddress string, _portNumber string) *Room {

	return &Room{description: _description}
}
