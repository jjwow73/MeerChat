package client

import (
	"../message"

	"bufio"
	"fmt"
	"golang.org/x/net/websocket"
	"math/rand"
	"os"
	"time"
)

func Client(port *string) (*websocket.Conn, error) {
	return websocket.Dial(fmt.Sprintf("ws://localhost:%s", *port), "", mockedIP())
}

func ReceiveMessage(ws *websocket.Conn) {
	var m message.Message
	for {
		err := websocket.JSON.Receive(ws, &m)
		if err != nil {
			fmt.Println("Error receiving message: ", err.Error())
			break
		}
		fmt.Println("Message: ", m)
	}
}

func SendMessage(ws *websocket.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		m := message.Message{
			Text: text,
		}
		err := websocket.JSON.Send(ws, m)
		if err != nil {
			fmt.Println("Error sending message: ", err.Error())
			break
		}
	}
}

func mockedIP() string {
	var arr [4]int
	for i := 0; i < 4; i++ {
		rand.Seed(time.Now().UnixNano())
		arr[i] = rand.Intn(256)
	}
	return fmt.Sprintf("http://%d.%d.%d.%d", arr[0], arr[1], arr[2], arr[3])
}
