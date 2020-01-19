package client

import (
	"bufio"
	"log"
	"os"
)

const (
	commandRegex        = "meer [a-z0-9]+"
	commandJoinRegex    = "meer join [0-9]+(?:\\.[0-9]+){3}:[0-9]+ [a-z0-9]+ [a-z0-9]+"
	commandLeaveRegex   = "meer leave [a-z0-9]+"
	commandEnterRegex   = "meer room [a-z0-9]+"
	commandMessageRegex = "meer message (?P<message>.+)"
	commandListRegex    = "meer list"
)

func init() {
	scanner = bufio.NewScanner(os.Stdin)
}

var scanner *bufio.Scanner

func handleInput() {
	for scanner.Scan() {
		text := scanner.Text()
		switch {
		case ifCommand(text):
			handleCommand(text)
		default:
			log.Println("not command")
		}
	}
}
