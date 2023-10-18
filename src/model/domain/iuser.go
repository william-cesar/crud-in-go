package domain

import "github.com/william-cesar/crud-in-go/src/config/ierrors"

type IUser interface {
	GetId() string
	GetName() string
	GetEmail() string
	GetPassword() string
	GetAge() int8
	IsActive() bool
	EncryptPassword()
	SendActivationEmail(email, id string) *ierrors.TError
	SetID(id string)
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
