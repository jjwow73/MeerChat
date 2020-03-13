package params

type JoinArgs struct {
	IP           string
	Port         string
	RoomId       string
	RoomPassword string
}

type SendArgs struct {
	Message string
}

type Reply struct{}
