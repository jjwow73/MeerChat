package model

import (
	"github.com/gorilla/websocket"
	"github.com/jjwow73/MeerChat/pkg/chat"
	"github.com/jjwow73/MeerChat/pkg/params"
	"log"
	"net/url"
)

type Room struct {
	id       string
	password string
	ip       string
	port     string
	conn     *websocket.Conn
}

func NewRoom(args *params.JoinArgs, username string) (*Room, error) {
	// TODO: input 검증

	conn, err := join(args.RoomId, args.RoomPassword, args.IP, args.Port, username)
	if err != nil {
		return nil, err
	}

	return &Room{
		id:       args.RoomId,
		password: args.RoomPassword,
		ip:       args.IP,
		port:     args.Port,
		conn:     conn,
	}, nil
}

func join(id, password, ip, port, username string) (*websocket.Conn, error) {
	query := "id=" + id + "&password=" + password + "&name=" + username
	u := url.URL{Scheme: "ws", Host: ip + ":" + port, Path: "/ws", RawQuery: query}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	return conn, err
}

func (r *Room) listenAndSendTo(ch chan *chat.MessageProtocol) {
	for {
		_, messageProtocolReceived, err := r.conn.ReadMessage()
		if err != nil {
			select {
			case <-ch:
			default:
				close(ch)
			}
			return
		}
		messageProtocol := &chat.MessageProtocol{}
		if err := messageProtocol.Unmarshal(messageProtocolReceived); err != nil {
			log.Println("json parsing:", err)
			continue
		}
		ch <- messageProtocol
	}
}

func (r *Room) send(message string) error {
	err := r.conn.WriteMessage(websocket.TextMessage, []byte(message))
	// TODO: RoomManager에서 error가 있다면 closeRoom channel
	if err != nil {
		return err
	}
	return nil
}

func (r *Room) closeRoom() error {
	err := r.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	return err
}
