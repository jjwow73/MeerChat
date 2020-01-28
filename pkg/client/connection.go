package client

import (
	"../chat"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

type connection struct {
	addr     string
	id       string
	password string
	conn     *websocket.Conn
	done     chan interface{}
}

type connectionMessage struct {
	c           *connection
	jsonMessage chat.Message
}

func newConnection(addr, id, password, name string) (c *connection, err error) {
	conn, err := connectToWebsocket(addr, id, password, name)
	if err != nil {
		log.Println("dial:", err)
		return c, err
	}
	return &connection{
		addr:     addr,
		id:       id,
		password: password,
		conn:     conn,
		done:     make(chan interface{}),
	}, nil
}

func (c *connection) readMessage(channel chan *connectionMessage) {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			select {
			case <-c.done: // normal closed
			default: // abnormal closed
				log.Println("connection", c.id, " read error:", err)
				close(c.done)
			}
			return
		}
		jsonMessage := chat.Message{}
		err = json.Unmarshal(message, &jsonMessage)
		if err != nil {
			log.Println("json parsing:", err)
			continue
		}
		channel <- &connectionMessage{c: c, jsonMessage: jsonMessage}
	}
}

func (c *connection) writeMessage(message string) {
	err := c.conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("write error:", err)
		close(c.done)
	}
}

func (c *connection) toString() string {
	return c.addr + " " + c.id
}

func connectToWebsocket(addr, id, password, name string) (conn *websocket.Conn, err error) {
	query := "id=" + id + "&" + "password=" + password + "&" + "name=" + name
	u := url.URL{Scheme: "ws", Host: addr, Path: "/ws", RawQuery: query}
	log.Printf("connecting to %s", u.String())

	conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	return conn, err
}
