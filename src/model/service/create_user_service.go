package service

import (
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	d "github.com/william-cesar/crud-in-go/src/model/domain"
)

func (us *tUserService) CreateUserService(
	newUser d.IUser,
) (d.IUser, *ierrors.TError) {
	newUser.EncryptPassword()
	savedUser, err := us.repository.CreateUser(newUser)

	if err != nil {
		return nil, err
	}

	return savedUser, nil
}