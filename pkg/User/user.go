package user

import (
	"fmt"
	"github.com/wkd3475/MeerChat/conf"
	"github.com/wkd3475/MeerChat/pkg/util"
)

type User struct {
	name       string
	ipAddress  string
	portNumber string
}

func NewUser() *User {
	_ip := util.GetInboundIP()
	conf.Ip = _ip
	return &User{name: conf.Nickname, ipAddress: _ip, portNumber: conf.Port}
}

type UserError struct {
	OriginalError error
}

func (userErr *UserError) Error() string {
	return fmt.Sprintf("User error : %v", userErr.OriginalError)
}

func (user *User) GetUserInfo() (string, string, string) {
	return user.name, user.ipAddress, user.portNumber
}

func (user *User) ModifyNickname(nickname string) error {
	if len(nickname) < 0 || len(nickname) > 8 {
		err := fmt.Errorf("nickname length -> 0~8")
		return &UserError{OriginalError: err}
	}
	user.name = nickname
	return nil
}