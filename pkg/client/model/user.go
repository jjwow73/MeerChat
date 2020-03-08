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
	user.name = name
}
