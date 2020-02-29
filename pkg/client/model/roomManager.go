package model

type RoomManager struct {
	roomToConnection map[*Room]*Connection
	focusedRoom      *Room
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		roomToConnection: make(map[*Room]*Connection),
		focusedRoom:      nil,
	}
}
