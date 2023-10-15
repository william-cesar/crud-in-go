package domain

type IUser interface {
	GetId() string
	GetName() string
	GetEmail() string
	GetPassword() string
	GetAge() int8
	IsActive() bool
}

func NewUser(
	email, password, name string,
	age int8,
) IUser {
	return &tUser{
		email:    email,
		password: password,
		name:     name,
		age:      age,
		active:   false,
	}
}

func NewUserLogin(
	email, password string,
) IUser {
	return &tUser{
		email:    email,
		password: password,
	}
}
