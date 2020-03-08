package model

import (
	"github.com/gorilla/websocket"
	"github.com/jjwow73/MeerChat/pkg/chat"
	"log"
	"net/url"
)

type Connection struct {
	conn *websocket.Conn
	ch   chan *chat.MessageProtocol
}

func NewConnection(room Room, user User) *Connection {
	conn, err := join(room, user)
	if err != nil {
		log.Fatal(err)
	}

	return &Connection{
		conn: conn,
		ch:   make(chan *chat.MessageProtocol),
	}
}

func join(room Room, user User) (*websocket.Conn, error) {
	query := "id=" + room.id + "&password=" + room.password + "&name=" + user.name
	u := url.URL{Scheme: "ws", Host: room.ip + ":" + room.port, Path: "/ws", RawQuery: query}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	return conn, err
}

func (c *Connection) listener() {
	for {
		_, messageProtocolReceived, err := c.conn.ReadMessage()
		if err != nil {
			select {
			case <-c.ch:
			default:
				close(c.ch)
			}
			return
		}
		messageProtocol := &chat.MessageProtocol{}
		if err := messageProtocol.Unmarshal(messageProtocolReceived); err != nil {
			log.Println("json parsing:", err)
			continue
		}
		c.ch <- messageProtocol
	}
}

func (c *Connection) send(message string) {
	err := c.conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		close(c.ch)
	}
}

func (c *Connection) close() {
	c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	close(c.ch)
}
