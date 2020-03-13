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

type LeaveArgs struct {
	IP     string
	Port   string
	RoomId string
}

type ListArgs struct {
}

type Reply struct{}
