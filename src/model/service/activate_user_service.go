package service

import (
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
)

func (us *tUserService) ActivateUserService(id string) *ierrors.TError {
	err := us.repository.ActivateUser(id)

	return err
}
