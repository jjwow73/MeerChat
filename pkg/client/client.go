package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func StartClient(targetAddress string) {
	conn, err := net.Dial("tcp", targetAddress)
	if nil != err {
		log.Printf("failed to connect to server %s", targetAddress)
		return
	}

	fmt.Println("input -1 if you want to exit")
	fmt.Println("----------------------------")

	for {
		input := bufio.NewReader(os.Stdin)
		message, err := input.ReadString('\n')
		if err != nil {
			log.Printf("fail to read input message; err: %v", err)
		}

		if message == "-1\n" {
			conn.Close()
			break
		}
		conn.Write([]byte(message))
	}
}