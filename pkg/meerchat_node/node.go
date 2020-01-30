package meerchat_node

import (
	"bytes"
	"github.com/wkd3475/MeerChat/pkg/protocol"
	"github.com/wkd3475/MeerChat/pkg/room"
	"github.com/wkd3475/MeerChat/pkg/user"
)

type Node struct {
	User *user.User
	Room []room.Room
}

func NewNode() *Node {
	_user := user.NewUser()
	return &Node{User: _user}
}

