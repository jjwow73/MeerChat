package client

import (
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

type admin struct {
	connections       map[string]*connection
	outputChannel     chan *connectionMessage
	focusedConnection *connection
}

func newAdmin() *admin {
	return &admin{
		connections:   make(map[string]*connection),
		outputChannel: make(chan *connectionMessage),
	}
}

var a *admin

func init() {
	a = newAdmin()
}

func Start() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go a.printMessageOfFocusedConnection()
	go readInputs(a)

	<-interrupt
	for id := range a.connections {
		a.leaveConnection(id)
	}

}

func (a *admin) joinConnection(addr, id, password, name string) {
	connection, err := newConnection(addr, id, password, name)
	if err != nil {
		return
	}
	a.connections[id] = connection
	a.focusConnection(id)
	log.Println("join to connection", id)
	go connection.readMessage(a.outputChannel)
	go a.deferRemoveConnection(connection)
}

func (a *admin) leaveConnection(id string) {
	connection, exist := a.getConnection(id)
	if !exist {
		return
	}
	log.Println("leave connection", id)
	connection.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	close(connection.done)
}

func (a *admin) focusConnection(id string) {
	c, exist := a.getConnection(id)
	if !exist {
		return
	}
	log.Println("focus to the connection", id)
	a.focusedConnection = c
}

func (a *admin) writeMessageInFocusedConnection(message string) {
	if a.focusedConnection == nil {
		log.Println("there aren't focused connection")
		return
	}
	a.focusedConnection.writeMessage(message)
}

func (a *admin) getConnectionList() {
	log.Println("get list")
	for idx, c := range a.connections {
		log.Println(idx + ". " + c.toString())
	}
}

func (a *admin) getConnection(id string) (c *connection, exist bool) {
	c, exist = a.connections[id]
	if !exist {
		log.Println("connection doesn't exist")
	}
	return
}

func (a *admin) printMessageOfFocusedConnection() {
	for {
		cm := <-a.outputChannel
		if a.focusedConnection == cm.c {
			log.Println(cm.jsonMessage.Name, ":", string(cm.jsonMessage.Content))
		}
	}
}

func (a *admin) deferRemoveConnection(c *connection) {
	<-c.done
	a.removeConnection(c)
}

func (a *admin) removeConnection(c *connection) {
	if err := c.conn.Close(); err != nil {
		log.Println(err)
	}
	if c == a.focusedConnection {
		a.focusedConnection = nil
	}
	delete(a.connections, c.id)
	c = nil
}
