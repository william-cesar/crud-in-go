package domain

import "github.com/william-cesar/crud-in-go/src/config/ierrors"

type IUser interface {
	GetId() string
	GetName() string
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetActive() bool
	IsActive() bool
	EncryptPassword()
	SendActivationEmail(email, id string) *ierrors.TError
	SetID(id string)
	SetActive()
}

func NewUser(
	email, password, name string,
	age int8,
	active bool,
) IUser {
	return &tUser{
		email:    email,
		password: password,
		name:     name,
		age:      age,
		active:   active,
	}
}

func NewUserUpdate(
	password, name string,
	age int8,
) IUser {
	return &tUser{
		password: password,
		name:     name,
		age:      age,
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
