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
	NumberOfRoom int
	ActiveRoom int
}

func NewNode() *Node {
	_user := user.NewUser()
	return &Node{User: _user, NumberOfRoom: 0, ActiveRoom: -1}
}

func (node *Node) CommandReceiver(cuiChan chan protocol.Command) {
	for {
		select {
		case v := <-cuiChan:
			if v.OpCode == protocol.GetUserInfoCode {
				name, ip, port := node.User.GetUserInfo()
				answer := new(bytes.Buffer)
				answer.WriteString(name)
				answer.WriteString(ip)
				answer.WriteString(port)
				v.Answer = *answer
				cuiChan <- v
			} else if v.OpCode == protocol.ModifyNicknameCode {

			}
		}
	}
}

func (node *Node) CommandSender(command protocol.Command) {

}