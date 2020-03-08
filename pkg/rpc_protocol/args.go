package rpc_protocol

type Args struct {
	IP           string
	Port         string
	RoomId       string
	RoomPassword string
	UserName     string
	Message      string
}

type Reply struct{}
