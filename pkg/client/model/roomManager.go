package model

import (
	"github.com/jjwow73/MeerChat/pkg/chat"
	"github.com/jjwow73/MeerChat/pkg/params"
	"log"
)

type RoomManager struct {
	roomsToChan map[*Room]chan *chat.MessageProtocol
	focusedRoom *Room
	outputChan  chan *chat.MessageProtocol
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
		log.Println(err)
		return
	}
	ch := make(chan *chat.MessageProtocol)
	rm.roomsToChan[room] = ch
	go room.listenAndSendTo(ch)
	go rm.listen(room)
}

func (rm *RoomManager) listen(room *Room) {
	for {
		select {
		case message, ok := <-rm.roomsToChan[room]:
			if !ok {
				rm.Delete(room)
				return
			}

			if room == rm.focusedRoom {
				rm.outputChan <- message
			}
		}
	}
}

func (rm *RoomManager) Send(args *params.SendArgs, username string) {
	//TODO: username 컨트롤러? RM?
	if err := rm.focusedRoom.send(args.Message); err != nil {
		rm.Delete(rm.focusedRoom)
	}
}

func (rm *RoomManager) Delete(room *Room) {
	//TODO: 세분화된 delete
	rm.freeIfFocusedRoom(room)

	room.closeRoom()
	close(rm.roomsToChan[room])
	delete(rm.roomsToChan, room)
}

func (rm *RoomManager) freeIfFocusedRoom(room *Room) {
	if rm.focusedRoom == room {
		rm.SetFocusedRoom(nil)
	}
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
