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

func (rm *RoomManager) add(room *Room, connection *Connection) {
	rm.roomToConnection[room] = connection
}

func (rm *RoomManager) delete(room *Room) {
	rm.roomToConnection[room].close()
	delete(rm.roomToConnection, room)
}
