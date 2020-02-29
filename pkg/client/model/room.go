package model

type Room struct {
	id       string
	password string
	ip       string
	port     string
}

func NewRoom(id, password, ip, port string) (*Room, error) {
	// TODO: input 검증

	return &Room{id: id, password: password, ip: ip, port: port}, nil
}