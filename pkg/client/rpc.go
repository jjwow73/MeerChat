package client

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	Addr         string
	RoomId       string
	RoomPassword string
	ClientName   string
	Message      string
}

type Reply struct{}

type Intermediate struct{}

func (i *Intermediate) Join(args *Args, reply *Reply) error {
	a.joinConnection(args.Addr, args.RoomId, args.RoomPassword, args.ClientName)
	return nil
}

func (i *Intermediate) Leave(args *Args, reply *Reply) error {
	a.leaveConnection(args.RoomId)
	return nil
}

func (i *Intermediate) Focus(args *Args, reply *Reply) error {
	a.focusConnection(args.RoomId)
	return nil
}

func (i *Intermediate) Send(args *Args, reply *Reply) error {
	a.writeMessageInFocusedConnection(args.Message)
	return nil
}

func (i *Intermediate) List(args *Args, reply *Reply) error {
	a.getConnectionList()
	return nil
}

func RpcStart() {
	i := new(Intermediate)
	rpc.Register(i)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":12039")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
