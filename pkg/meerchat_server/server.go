package server

import (
	"fmt"
	"io"
	"log"
	"net"
)

func MakeServer(myPort int) {
	clientNum := 0
	portAddress := fmt.Sprintf(":%d", myPort)
	l, err := net.Listen("tcp", portAddress)
	if nil != err {
		log.Fatalf("fail to bind address to %d; err: %v", myPort, err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if nil != err {
			log.Printf("fail to accept; err: %v", err)
			continue
		}

		if clientNum == 0 {
			go ConnHandler(conn, clientNum)
			clientNum++
		} else {

			conn.Close()
		}

	}
}

func ConnHandler(conn net.Conn, clientNum int) {
	if addr, ok := conn.RemoteAddr().(*net.TCPAddr); ok {
		fmt.Println()
		fmt.Println("notice!", addr.IP.String(), "connected to you")
	}
	recvBuf := make([]byte, 4096)
	for {
		n, err := conn.Read(recvBuf)
		if nil != err {
			if io.EOF == err {
				log.Printf("connection is closed from client; %v", conn.RemoteAddr().String())
				clientNum--
				return
			}
			log.Printf("fail to receive data; err: %v", err)
			return
		}
		fmt.Print("receive : ")
		if 0 < n {
			data := recvBuf[:n]
			fmt.Print(string(data))
		}
	}
}