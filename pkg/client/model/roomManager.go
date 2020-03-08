package model

import (
	"github.com/jjwow73/MeerChat/pkg/chat"
	"github.com/jjwow73/MeerChat/pkg/params"
	"log"
)

type RoomManager struct {
	roomsToChan map[*Room]chan *chat.MessageProtocol
	focusedRoom *Room
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		roomsToChan: make(map[*Room]chan *chat.MessageProtocol),
		focusedRoom: nil,
	}
}

func (rm *RoomManager) Join(args *params.JoinArgs, username string) {
	room, err := NewRoom(args, username)
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan *chat.MessageProtocol)
	rm.roomsToChan[room] = ch
	go room.listener(ch)

	// TODO: RoomManager의 listener 구현, 아래와 같은 느낌
	//for {
	//	select {
	//	case message := <-ch:
	//		if room == rm.focusedRoom {
	//			rm.outputChan <- message
	//		}
	//	}
	//}
}

func (rm *RoomManager) Delete(room *Room) {
	if rm.focusedRoom == room {
		rm.SetFocusedRoom(nil)
	}

	room.closeRoom()
	close(rm.roomsToChan[room])
	delete(rm.roomsToChan, room)
}

func (rm *RoomManager) SetFocusedRoom(room *Room) {
	rm.focusedRoom = room
}

func (rm *RoomManager) GetRoomList() []*Room {
	rooms := make([]*Room, 0, len(rm.roomsToChan))
	for room := range rm.roomsToChan {
		rooms = append(rooms, room)
	}

	return rooms
}
