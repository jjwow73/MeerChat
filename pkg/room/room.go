package room

import (
	"github.com/wkd3475/MeerChat/pkg/user"
	"fmt"
	"github.com/wkd3475/MeerChat/pkg/util"
	"time"
)

const (
	public = 0
	private = 1
	defaultMaxUserNum = 20
)

type Room struct {
	name          string
	admin         user.User
	roomType	  int
	password      string
	description   string
	ipAddress     string
	portNumber    string
	maxUserNumber int32
	userList      []user.User
	blackList     []user.User
	updatedAt     int64
}

type roomError struct {
	msg string
}

func (roomErr *roomError) Error() string {
	return fmt.Sprintf("room error : %v", roomErr.msg)
}

func NewRoom(_name string, _admin user.User, _description string, _password string, _ipAddress string, _portNumber string) (*Room, error) {
	var _roomType int
	var hashedPassword string
	var err error
	var _userList []user.User
	_userList = append(_userList, _admin)

	if len(_name) < 0 || len(_name) > 10 {
		msg := "creat room error : name(<=10)"
		return nil, &roomError{msg}
	}

	if len(_description) < 0 || len(_description) > 30 {
		msg := "creat room error : description(<=30)"
		return nil, &roomError{msg}
	}

	if len(_password) == 0 {
		_roomType = public
	} else if len(_password) >= 0 && len(_password) <= 8 {
		_roomType = private
		hashedPassword, err = util.Generate(_password)
		if err != nil {
			return nil, err
		}
	} else {
		msg := "creat room error : password(<=8)"
		return nil, &roomError{msg}
	}

	return &Room{name: _name,
		admin: _admin,
		roomType: _roomType,
		password: hashedPassword,
		description: _description,
		ipAddress: _ipAddress,
		portNumber: _portNumber,
		maxUserNumber: defaultMaxUserNum,
		userList: _userList,
		updatedAt: time.Now().Unix()}, nil
}

func (room *Room) GetName() string {
	return room.name
}