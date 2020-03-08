package controller

import (
	"github.com/jjwow73/MeerChat/pkg/client/model"
	"github.com/jjwow73/MeerChat/pkg/rpc_protocol"
	"log"
)

type RpcService struct{}

var (
	roomManager *model.RoomManager
	user *model.User
)

const defaultName = "Meer"

func init() {
	roomManager = model.NewRoomManager()
	user = model.NewUser(defaultName)
}

func (rs *RpcService) Join(args *rpc_protocol.Args, reply rpc_protocol.Reply) error {
	room, err := model.NewRoom(args.RoomId, args.RoomPassword, args.IP, args.Port)
	if err != nil {
		log.Fatal(err)
	}
	connection := model.NewConnection(*room, *user)
	roomManager.Add(room, connection)
	return nil
}