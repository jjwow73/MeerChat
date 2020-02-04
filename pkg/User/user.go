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

type userError struct {
	msg string
}

func (userErr *userError) Error() string {
	return fmt.Sprintf("User error : %s", userErr.msg)
}

func (user *User) GetUserInfo() (string, string, string) {
	return user.name, user.ipAddress, user.portNumber
}

func (user *User) ModifyNickname(nickname string) error {
	if len(nickname) < 0 || len(nickname) > 8 {
		err := fmt.Sprintf("nickname length -> 0~8")
		return &userError{msg: err}
	}
	user.name = nickname
	return nil
}