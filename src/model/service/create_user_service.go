package service

import (
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	d "github.com/william-cesar/crud-in-go/src/model/domain"
)

func (us *tUserService) CreateUserService(
	newUser d.IUser,
) (d.IUser, *ierrors.TError) {
	user, err := us.repository.FindUserByEmail(newUser.GetEmail())

	if err != nil && err.StatusCode != 404 {
		return nil, err
	}

	if user != nil {
		return nil, ierrors.NewBadRequestError("User already exists.")
	}

	newUser.EncryptPassword()
	savedUser, err := us.repository.CreateUser(newUser)

	if err != nil {
		return nil, err
	}

	if err := savedUser.SendActivationEmail(
		savedUser.GetEmail(),
		savedUser.GetId(),
	); err != nil {
		if dErr := us.repository.DeleteUser(savedUser.GetId()); dErr != nil {
			return nil, dErr
		}
		return nil, err
	}

	return savedUser, nil
}
