package client

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

const (
	commandRegex        = "meer [a-z0-9]+"
	commandJoinRegex    = "meer join (?P<addr>[0-9]+(?:\\.[0-9]+){3}:[0-9]+) (?P<id>[a-z0-9]+) (?P<password>[a-z0-9]+)"
	commandLeaveRegex   = "meer leave (?P<id>[a-z0-9]+)"
	commandEnterRegex   = "meer room ([?P<id>a-z0-9]+)"
	commandMessageRegex = "meer message (?P<message>.+)"
	commandListRegex    = "meer list"
)

func handleInput() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		switch {
		case matchWithPattern(commandRegex, text):
			handleCommand(text)
		default:
			log.Println("not command")
		}
	}
}

func handleCommand(text string) {
	switch {
	case matchWithPattern(commandJoinRegex, text):
		handleCommandJoin(text)
	case matchWithPattern(commandLeaveRegex, text):
		handleCommandLeave(text)
	case matchWithPattern(commandEnterRegex, text):
		handleCommandEnter(text)
	case matchWithPattern(commandMessageRegex, text):
		handleCommandMessage(text)
	case matchWithPattern(commandListRegex, text):
		handleCommandList()
	default:
	}
}

func handleCommandJoin(text string) {
	parsedText := getParsedText(commandJoinRegex, text)
	joinRoom(parsedText[1], parsedText[2], parsedText[3])
}

func handleCommandLeave(text string) {
	parsedText := getParsedText(commandLeaveRegex, text)
	leaveRoom(parsedText[1])
}

func handleCommandEnter(text string) {
	parsedText := getParsedText(commandEnterRegex, text)
	focusRoom(parsedText[1])
}

func handleCommandMessage(text string) {
	parsedText := getParsedText(commandMessageRegex, text)
	writeMessage(parsedText[1])
}

func handleCommandList() {
	getRoomList()
}

func matchWithPattern(pattern string, text string) bool {
	match, err := regexp.MatchString(pattern, text)
	if err != nil {
		log.Fatal(err)
	}
	return match
}

func getParsedText(pattern string, text string) []string {
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}
	return re.FindStringSubmatch(text)
}
