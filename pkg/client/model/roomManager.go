package model

import (
	"errors"
	"fmt"
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
	rm.SetFocusedRoom(room)
	go room.listenAndSendTo(ch)
	go rm.listen(room)
}

func (rm *RoomManager) listen(room *Room) {
	for {
		select {
		case message, ok := <-rm.roomsToChan[room]:
			if !ok {
				rm.delete(room)
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
		rm.delete(rm.focusedRoom)
	}
}

func (rm *RoomManager) delete(room *Room) {
	rm.freeIfFocusedRoom(room)
	rm.close(room)
	rm.closeChan(room)
	rm.removeRoomsToChan(room)
}

func (rm *RoomManager) freeIfFocusedRoom(room *Room) {
	if rm.focusedRoom == room {
		rm.SetFocusedRoom(nil)
	}
}

func (rm *RoomManager) close(room *Room) {
	if err := room.closeRoom(); err != nil {
		log.Println("이미 닫힌 room을 닫으려 시도")
	}
}

func (rm *RoomManager) closeChan(room *Room) {
	select {
	case <-rm.roomsToChan[room]:
		close(rm.roomsToChan[room])
	default:
		log.Print("이미 닫힌 chan을 닫으려 시도")
	}
}

func (rm *RoomManager) removeRoomsToChan(room *Room) {
	if _, ok := rm.roomsToChan[room]; ok {
		delete(rm.roomsToChan, room)
	} else {
		log.Println("이미 삭제된 map 제거 시도")
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

func (rm *RoomManager) Leave(args *params.LeaveArgs) {
	room, err := rm.findRoom(args.IP, args.Port, args.RoomId)
	if err != nil {
		log.Println(err)
		return
	}
	rm.delete(room)
}

func (rm *RoomManager) findRoom(ip, port, roomId string) (*Room, error) {
	for room, _ := range rm.roomsToChan {
		if (room.ip == ip) && (room.port == port) && (room.id == roomId) {
			return room, nil
		}
	}
	return nil, errors.New("no such room")
}

func (rm *RoomManager) Focus(args *params.FocusArgs) {
	room, err := rm.findRoom(args.IP, args.Port, args.RoomId)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("현재 room", room, " 주목 room", rm.focusedRoom)
	rm.SetFocusedRoom(room)
	fmt.Println("바뀐 뒤, 현재 room", room, " 주목 room", rm.focusedRoom)
}
