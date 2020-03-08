package rpc_protocol

type Args struct {
	Addr         string
	RoomId       string
	RoomPassword string
	UserName     string
	Message      string
}

type Reply struct{}