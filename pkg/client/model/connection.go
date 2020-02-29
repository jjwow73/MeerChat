package model

import (
	"github.com/gorilla/websocket"
	"github.com/jjwow73/MeerChat/pkg/chat"
	"log"
	"net/url"
)

type connection struct {
	conn *websocket.Conn
	ch   chan *chat.MessageProtocol
}

func (c *connection) join(room *Room, user *User) (*websocket.Conn, error) {
	query := "id=" + room.id + "&password=" + room.password + "&name=" + user.name
	u := url.URL{Scheme: "ws", Host: room.ip + ":" + room.port, Path: "/ws", RawQuery: query}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	return conn, err
}

func (c *connection) listener() {
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

func (c *connection) send(message string) {
	err := c.conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		close(c.ch)
	}
}

func (c *connection) close() {
	c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	close(c.ch)
}
