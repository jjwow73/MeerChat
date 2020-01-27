package main

import (
	"flag"
	"fmt"
	"github.com/wkd3475/MeerChat/pkg/client"
	"github.com/wkd3475/MeerChat/pkg/server"
)

func main() {
	myPort := flag.Int("p", 7000, "port")
	var targetAddress string
	flag.Parse()

	go server.MakeServer(*myPort)
	fmt.Println("My server started")

	for {
		var input int
		checker := false
		fmt.Println("1. set server")
		fmt.Println("0. exit")
		fmt.Scanf("%d", &input)

		switch input {
		case 0:
			checker = true
		case 1:
			fmt.Printf("tartget server : ")
			fmt.Scanf("%s", &targetAddress)
			client.StartClient(targetAddress)
		default:
			fmt.Println("wrong input")
		}

		if checker {
			break
		}
	}
}
