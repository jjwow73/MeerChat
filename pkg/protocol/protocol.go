package protocol

import "bytes"

const (
	GetUserInfoCode = 100
	ModifyNicknameCode = 101
)

type Command struct {
	CommandType int
	Options []byte
	Answer bytes.Buffer
}

