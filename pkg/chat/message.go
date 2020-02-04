package chat

type Message struct {
	Content []byte `json:"content"`
	Name    string `json:"name"`
}
