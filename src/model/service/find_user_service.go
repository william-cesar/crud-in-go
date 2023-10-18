package service

import (
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	d "github.com/william-cesar/crud-in-go/src/model/domain"
)

func (us *tUserService) FindUserService(email string) (d.IUser, *ierrors.TError) {
	user, err := us.repository.FindUser(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}
