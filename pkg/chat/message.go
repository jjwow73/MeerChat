package chat

import (
	"encoding/json"
)

type MessageProtocol struct {
	Message []byte `json:"message"`
	Name    string `json:"name"`
}

func (m *MessageProtocol) Unmarshal(messageProtocol []byte) error {
	err := json.Unmarshal(messageProtocol, m)
	return err
}
