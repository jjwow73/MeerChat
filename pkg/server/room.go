package server

import "log"

var rooms roomList

func init() {
	rooms = roomList{}
	rooms.rooms = map[string]*Room{}
}

type Room struct {
	id       string
	password string
	Hub      *Hub
}

type roomList struct {
	rooms map[string]*Room
}

func GetRoom(id string, password string) (room *Room, auth bool) {
	room, exist := rooms.rooms[id]
	if exist {
		if password == room.password {
			return room, true
		}
		return room, false
	}
	hub := NewHub()
	go hub.Run()
	room = &Room{id: id, password: password, Hub: hub}
	rooms.rooms[id] = room
	return room, true
}

func RemoveRoom(id string) (exist bool){
	room, exist := rooms.rooms[id]
	if !exist {
		return exist
	}
	log.Println("delete room :", room.id)
	room = nil
	delete(rooms.rooms, id)
	return exist
}

