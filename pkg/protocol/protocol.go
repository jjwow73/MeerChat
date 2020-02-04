package protocol

import "bytes"

const (
	GetUserInfoCode = byte(0x10)
	ModifyNicknameCode = byte(0x11)
	UpdateRoomInfoCode = byte(0x20)
)

type Command struct {
	OpCode byte
	Options []byte
	Answer bytes.Buffer
}

func (command *Command) Command2Bytes() ([]byte, error) {
	msg := make([]byte, 4096)
	var err error

	if command.OpCode == UpdateRoomInfoCode {

		msg[0] = UpdateRoomInfoCode
	}

	return msg, err
}