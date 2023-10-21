package service

import (
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/model/domain"
)

func (us *tUserService) LoginUserService(user domain.IUser) (string, *ierrors.TError) {
	user.EncryptPassword()

	if err := us.findUserByEmailAndPassword(user); err != nil {
		return "", err
	}

	return user.GenerateToken()
}
