package controller

import (
	"github.com/jjwow73/MeerChat/pkg/client/model"
	"github.com/jjwow73/MeerChat/pkg/params"
)

type RpcService struct{}

var (
	roomManager *model.RoomManager
	user        *model.User
)

const defaultName = "Meer"

func init() {
	roomManager = model.NewRoomManager()
	user = model.NewUser(defaultName)
}

func (rs *RpcService) Join(args *params.JoinArgs, reply params.Reply) error {
	roomManager.Join(args, user.GetUserName())
	return nil
}
