package service

import (
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/model/domain"
)

func (us *tUserService) FindUserByEmailService(email string) (domain.IUser, *ierrors.TError) {
	user, err := us.repository.FindUserByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *tUserService) FindUserByIdService(id string) (domain.IUser, *ierrors.TError) {
	user, err := us.repository.FindUserById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
