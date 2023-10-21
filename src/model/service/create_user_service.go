package service

import (
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/config/logger"
	"github.com/william-cesar/crud-in-go/src/model/domain"
)

func (us *tUserService) CreateUserService(
	newUser domain.IUser,
) (domain.IUser, *ierrors.TError) {
	logger.NewInfoLog(logger.JOURNEY["CREATE_SERVICE"], logger.MESSAGE["INIT"]["CREATION"])

	user, err := us.repository.FindUserByEmail(newUser.GetEmail())

	if err != nil && err.StatusCode != 404 {
		logger.NewErrorLog(logger.JOURNEY["CREATE_SERVICE"], logger.MESSAGE["ERROR"]["NO_USER"])
		return nil, err
	}

	if user != nil {
		logger.NewErrorLog(logger.JOURNEY["CREATE_SERVICE"], logger.MESSAGE["ERROR"]["UNIQUE"])
		return nil, ierrors.NewBadRequestError("User already exists.")
	}

	newUser.EncryptPassword()
	savedUser, err := us.repository.CreateUser(newUser)

	if err != nil {
		logger.NewErrorLog(logger.JOURNEY["CREATE_SERVICE"], logger.MESSAGE["ERROR"]["CREATION"])
		return nil, err
	}

	if err := savedUser.SendActivationEmail(
		savedUser.GetEmail(),
		savedUser.GetId(),
	); err != nil {
		logger.NewErrorLog(logger.JOURNEY["CREATE_SERVICE"], logger.MESSAGE["ERROR"]["SEND_EMAIL"])
		logger.NewErrorLog(logger.JOURNEY["CREATE_SERVICE"], logger.MESSAGE["ERROR"]["CREATION"])

		if dErr := us.repository.DeleteUser(savedUser.GetId()); dErr != nil {
			return nil, dErr
		}
		return nil, err
	}

	logger.NewInfoLog(logger.JOURNEY["CREATE_SERVICE"], logger.MESSAGE["OK"]["CREATED"])
	return savedUser, nil
}
