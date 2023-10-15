package domain

import (
	"crypto/sha256"
	"encoding/hex"
)

type tUser struct {
	id       string
	name     string
	email    string
	password string
	age      int8
	active   bool
}

func (u *tUser) GetId() string {
	return u.id
}

func (u *tUser) GetName() string {
	return u.name
}

func (u *tUser) GetEmail() string {
	return u.email
}

func (u *tUser) GetPassword() string {
	return u.password
}

func (u *tUser) GetAge() int8 {
	return u.age
}

func (u *tUser) IsActive() bool {
	return u.active
}

func (u *tUser) EncryptPassword() {
	hash := sha256.New()
	defer hash.Reset()

	hash.Write([]byte(u.password))
	u.password = hex.EncodeToString(hash.Sum(nil))
}
