package client

import (
	"log"
	"os"
	"os/signal"

	"../chat"
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
	log.Println("jjong:ADMIN START")
	a = newAdmin()
	go a.printAllMessageFromOutputChannel()
}

func Start() {
	log.Println("jjong:START")
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
	//log.Println("join to connection", id)
	// connAddr, connId, connNickname := a.focusedConnection.GetConnInfo()
	//log.Printf("jjong: JOIN %v, %v, %v", connAddr, connId, connNickname)
	go connection.readMessage(a.outputChannel)
	//server로 outputChannel로 보내기.
	jsonMessage := chat.Message{Content: []byte("Join to connection:" + id), Name: "Local"}
	a.outputChannel <- &connectionMessage{c: connection, jsonMessage: jsonMessage}

	go a.deferRemoveConnection(connection)

}

func (a *admin) leaveConnection(id string) {
	connection, exist := a.getConnection(id)
	if !exist {
		jsonMessage := chat.Message{Content: []byte("Error leaving connection:" + id), Name: "Local"}
		a.outputChannel <- &connectionMessage{c: nil, jsonMessage: jsonMessage}
		return
	}

	jsonMessage := chat.Message{Content: []byte("leave connection" + id), Name: "Local"}
	a.outputChannel <- &connectionMessage{c: connection, jsonMessage: jsonMessage}

	// log.Println("leave connection", id)
	connection.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	close(connection.done)
}

func (a *admin) focusConnection(id string) {

	c, exist := a.getConnection(id)
	if !exist {
		return
	}

	jsonMessage := chat.Message{Content: []byte("focus to the connection: " + id), Name: "Local"}
	a.outputChannel <- &connectionMessage{c: c, jsonMessage: jsonMessage}

	a.focusedConnection = c

}

func (a *admin) writeMessageInFocusedConnection(message string) {
	if a.focusedConnection == nil {
		jsonMessage := chat.Message{Content: []byte("there aren't focused connection"), Name: "Local"}
		a.outputChannel <- &connectionMessage{c: nil, jsonMessage: jsonMessage}
		return
	}
	a.focusedConnection.writeMessage(message)
}

func (a *admin) getConnectionList() {
	// log.Println("get list")
	getList := "Get List\n"
	for idx, c := range a.connections {
		getList += idx + ". " + c.toString() + "\n"
		// log.Println(idx + ". " + c.toString())
	}
	jsonMessage := chat.Message{Content: []byte(getList), Name: "Local"}
	a.outputChannel <- &connectionMessage{c: nil, jsonMessage: jsonMessage}
}

func (a *admin) getConnection(id string) (c *connection, exist bool) {
	c, exist = a.connections[id]
	if !exist {
		jsonMessage := chat.Message{Content: []byte("connection " + id + "doesn't exist"), Name: "Local"}
		a.outputChannel <- &connectionMessage{c: nil, jsonMessage: jsonMessage}
	}
	return
}

func (a *admin) printMessageOfFocusedConnection() {
	//log.Println("printMessageOfConnection")
	for cm := range a.outputChannel {
		log.Println("PRINT:", cm.jsonMessage.Name, ":", string(cm.jsonMessage.Content))
		if a.focusedConnection == cm.c {
			log.Println(cm.jsonMessage.Name, ":", string(cm.jsonMessage.Content))
		}
	}
	/*
		for {
			cm := <-a.outputChannel
			log.Println("PRINT:", cm.jsonMessage.Name, ":", string(cm.jsonMessage.Content))
			if a.focusedConnection == cm.c {
				log.Println(cm.jsonMessage.Name, ":", string(cm.jsonMessage.Content))
			}
		}*/
}

func (a *admin) printAllMessageFromOutputChannel() {
	//log.Println("printAllMessageFromOutputChannel START")
	for cm := range a.outputChannel {
		log.Println(cm.jsonMessage.Name, ":", string(cm.jsonMessage.Content))
	}
}

func (a *admin) deferRemoveConnection(c *connection) {
	<-c.done
	a.removeConnection(c)
}

func (a *admin) removeConnection(c *connection) {
	if err := c.conn.Close(); err != nil {

		jsonMessage := chat.Message{Content: []byte("Connection Closed"), Name: "Server"}
		a.outputChannel <- &connectionMessage{c: c, jsonMessage: jsonMessage}

		// log.Println(err)
	}
	if c == a.focusedConnection {
		a.focusedConnection = nil
	}
	delete(a.connections, c.id)
	c = nil
}
