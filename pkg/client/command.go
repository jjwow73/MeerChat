package client

import (
	"regexp"
	"strings"
)

func ifCommand(text string) bool {
	ifCommand, _ := regexp.MatchString(commandRegex, text)
	return ifCommand
}

func handleCommand(text string) {
	handleCommandJoin(text)
	handleCommandLeave(text)
	handleCommandEnter(text)
	handleCommandMessage(text)
	handleCommandList(text)
}

func handleCommandJoin(text string) {
	ifCommandJoin, _ := regexp.MatchString(commandJoinRegex, text)
	if !ifCommandJoin {
		return
	}
	parseText := strings.Split(text, " ")
	joinRoom(parseText[2], parseText[3], parseText[4])
}

func handleCommandLeave(text string) {
	ifCommandLeave, _ := regexp.MatchString(commandLeaveRegex, text)
	if !ifCommandLeave {
		return
	}
	parseText := strings.Split(text, " ")
	leaveRoom(parseText[2])
}

func handleCommandEnter(text string) {
	ifCommandEnter, _ := regexp.MatchString(commandEnterRegex, text)
	if !ifCommandEnter {
		return
	}
	parseText := strings.Split(text, " ")
	focusRoom(parseText[2])
}

func handleCommandMessage(text string) {
	ifCommandMessage, _ := regexp.MatchString(commandMessageRegex, text)
	if !ifCommandMessage {
		return
	}
	parseText := strings.Split(text, " ")
	writeMessage(parseText[2])
}

func handleCommandList(text string) {
	ifCommandList, _ := regexp.MatchString(commandListRegex, text)
	if !ifCommandList {
		return
	}
	getRoomList()
}
