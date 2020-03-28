package controller

import (
	"fmt"
	"github.com/jjwow73/MeerChat/pkg/chat"
	"github.com/jjwow73/MeerChat/pkg/client/model"
	"github.com/jjwow73/MeerChat/pkg/client/view"
	"github.com/jjwow73/MeerChat/pkg/params"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
)

type RpcService struct{}

var (
	roomManager *model.RoomManager
	user        *model.User
	outputChan  chan *chat.MessageProtocol
)

const (
	defaultName = "Meer"
)

func init() {
	renderer := view.Renderer{}
	outputChan = make(chan *chat.MessageProtocol)
	roomManager = model.NewRoomManager(outputChan, renderer)
	user = model.NewUser(defaultName)
}

func (rs *RpcService) Join(args *params.JoinArgs, reply *params.Reply) error {
	roomManager.Join(args, user.GetUserName())
	return nil
}

func (rs *RpcService) Send(args *params.SendArgs, reply *params.Reply) error {
	roomManager.Send(args, user.GetUserName())
	return nil
}

func (rs *RpcService) Leave(args *params.LeaveArgs, reply *params.Reply) error {
	roomManager.Leave(args)
	return nil
}

func (rs *RpcService) List(args *params.ListArgs, reply *params.Reply) error {
	//TODO: roomlist 배열을 string으로 변환해서 outputchannel
	rooms := roomManager.GetRoomList()
	fmt.Println(rooms)
	return nil
}

func (rs *RpcService) Focus(args *params.FocusArgs, reply *params.Reply) error {
	roomManager.Focus(args)
	return nil
}

func (rs *RpcService) Name(args *params.NameArgs, reply *params.Reply) error {
	user.SetUserName(args.Name)
	view.PrintUserInfo(args.Name)
	return nil
}

func RpcStart() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	i := new(RpcService)
	rpc.Register(i)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":12039")
	defer l.Close()

	if e != nil {
		log.Fatal("listen error", e)
	}

	go http.Serve(l, nil)

	go func() {
		for {
			message := <-outputChan
			view.PrintChatMessage(message)
		}
	}()

	go view.Start()

	<-interrupt
}
