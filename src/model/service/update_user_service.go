package service

import (
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/model/domain"
)

func (us *tUserService) UpdateUserService(id string, user domain.IUser) *ierrors.TError {
	dbUser, err := us.repository.FindUserById(id)

	if err != nil {
		return err
	}

	if dbUser != nil && !dbUser.GetActive() {
		return ierrors.NewBadRequestError("User is not activated.")
	}

	if user.GetPassword() != "" {
		user.EncryptPassword()
	}

	if err := us.repository.UpdateUser(id, user); err != nil {
		return err
	}

	return nil
}
