package service

import (
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
)

func (us *tUserService) DeleteUserService(id string) *ierrors.TError {
	err := us.repository.DeleteUser(id)

	return err
}
