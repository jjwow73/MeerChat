package client

import (
	"github.com/gorilla/websocket"
	"log"
)

func init() {
	roomList = rooms{
		rooms: make(map[string]*room),
	}
}

type room struct {
	addr     string
	id       string
	password string
	conn     *websocket.Conn
	done     chan interface{}
}

type rooms struct {
	rooms     map[string]*room
	focusedId *string
}

var roomList rooms

func newRoom(addr string, id string, password string) (r *room, err error) {
	conn, err := connectToWebsocket(addr, id, password)
	if err != nil {
		log.Println("dial:", err)
		return r, err
	}
	return &room{
		addr:     addr,
		id:       id,
		password: password,
		conn:     conn,
		done:     make(chan interface{}),
	}, nil
}

func (r room) ifFocused() bool {
	return roomList.focusedId != nil && r.id == *roomList.focusedId
}

func getRoom(id string) (*room, bool) {
	room, exist := roomList.rooms[id]
	return room, exist
}

func joinRoom(addr string, id string, password string) {
	room, err := newRoom(addr, id, password)
	if err != nil {
		return
	}
	roomList.rooms[id] = room
	log.Println("join to room ", id)
	roomList.focusedId = &id
	go readMessage(room)
	go removeRoomAtTheEnd(room)
}

func getRoomList() {
	log.Println("get room list")
	for _, room := range roomList.rooms {
		log.Println("id: ", room.id, "addr: ", room.addr)
	}
}

func focusRoom(id string) {
	_, exist := getRoom(id)
	if !exist {
		log.Println("room doesn't exist")
		return
	}
	log.Println("enter to room ", id)
	roomList.focusedId = &id
}

func leaveRoom(id string) {
	room, exist := getRoom(id)
	if !exist {
		log.Println("room doesn't exist")
		return
	}
	log.Println("leave room ", id)
	room.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	close(room.done)
}

func removeRoomAtTheEnd(room *room) {
	<-room.done
	err := room.conn.Close()
	if err != nil {
		log.Println(err)
	}
	delete(roomList.rooms, room.id)
	if room.ifFocused() || len(roomList.rooms) == 0 {
		roomList.focusedId = nil
	}
	room = nil
}

func leaveAllRoom() {
	for id := range roomList.rooms {
		leaveRoom(id)
	}
}
