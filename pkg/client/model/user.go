package model

type User struct {
	name string
}

func NewUser(name string) *User {
	return &User{
		name: name,
	}
}

func (user User) GetUserName() string {
	return user.name
}

func (user *User) SetUserName(name string) {
	//TODO: server의 clientName 을 바꿔주는 통신
	user.name = name
}
