package params

type Args struct {
	IP           string
	Port         string
	RoomId       string
	RoomPassword string
	UserName     string
	Message      string
}

type JoinArgs struct {
	IP           string
	Port         string
	RoomId       string
	RoomPassword string
}

type Reply struct{}
