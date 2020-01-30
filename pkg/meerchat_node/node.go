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

func (node *Node) CommandReceiver(cuiChan chan protocol.Command) {
	for {
		select {
		case v := <-cuiChan:
			if v.CommandType == protocol.GetUserInfoCode {
				name, ip, port := node.User.GetUserInfo()
				answer := new(bytes.Buffer)
				answer.WriteString(name)
				answer.WriteString(ip)
				answer.WriteString(port)
				v.Answer = *answer
				cuiChan <- v
			} else if v.CommandType == protocol.ModifyNicknameCode {

			}
		}
	}
}
