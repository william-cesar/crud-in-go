package service

import (
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	d "github.com/william-cesar/crud-in-go/src/model/domain"
)

func (us *tUserService) UpdateUserService(id string, user d.IUser) *ierrors.TError {
	if user.GetPassword() != "" {
		user.EncryptPassword()
	}

	if err := us.repository.UpdateUser(id, user); err != nil {
		return err
	}

	return nil
}
