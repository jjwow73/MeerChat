package model

type User struct {
	name string
}

func (user User) getUserName() string {
	return user.name
}

func (user *User) setUserName(name string) {
	user.name = name
}