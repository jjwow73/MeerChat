package client

import "log"

const (
	commandRegex     = "meer [a-z0-9]+"
	commandJoinRegex = "meer join " +
		"(?P<addr>[0-9]+(?:\\.[0-9]+){3}:[0-9]+) " +
		"(?P<id>[a-z0-9]+) " +
		"(?P<password>[a-z0-9]+) " +
		"(?P<name>[a-z0-9]+)"
	commandLeaveRegex = "meer leave (?P<id>[a-z0-9]+)"
	commandFocusRegex = "meer focus (?P<id>[a-z0-9]+)"
	commandSendRegex  = "meer send (?P<chat>.+)"
	commandListRegex  = "meer list"
)

type commandJoin struct {
	addr     string
	id       string
	password string
	name     string
}

type commandLeave struct {
	id string
}

type commandFocus struct {
	id string
}

type commandSend struct {
	message string
}

type commandList struct{}

type commandUnknown struct{}

type command interface {
	handleCommand(a *admin)
}

func (c commandJoin) handleCommand(a *admin) {
	a.joinConnection(c.addr, c.id, c.password, c.name)
}

func (c commandLeave) handleCommand(a *admin) {
	a.leaveConnection(c.id)
}

func (c commandFocus) handleCommand(a *admin) {
	a.focusConnection(c.id)
}

func (c commandSend) handleCommand(a *admin) {
	a.writeMessageInFocusedConnection(c.message)
}

func (c commandList) handleCommand(a *admin) {
	a.getConnectionList()
}

func (c commandUnknown) handleCommand(a *admin) {
	log.Println("unknown command")
}
