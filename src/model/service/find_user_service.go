package service

import (
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/config/logger"
	"github.com/william-cesar/crud-in-go/src/model/domain"
)

func (us *tUserService) FindUserByEmailService(email string) (domain.IUser, *ierrors.TError) {
	logger.NewInfoLog(logger.JOURNEY["FIND_SERVICE"], logger.MESSAGE["INIT"]["FIND_EMAIL"])

	user, err := us.repository.FindUserByEmail(email)

	if err != nil {
		logger.NewErrorLog(logger.JOURNEY["FIND_SERVICE"], logger.MESSAGE["ERROR"]["NO_USER"])
		return nil, err
	}

	logger.NewInfoLog(logger.JOURNEY["FIND_SERVICE"], logger.MESSAGE["OK"]["FOUND"])
	return user, nil
}

func (us *tUserService) FindUserByIdService(id string) (domain.IUser, *ierrors.TError) {
	logger.NewInfoLog(logger.JOURNEY["FIND_SERVICE"], logger.MESSAGE["INIT"]["FIND_ID"])

	user, err := us.repository.FindUserById(id)

	if err != nil {
		logger.NewErrorLog(logger.JOURNEY["FIND_SERVICE"], logger.MESSAGE["ERROR"]["NO_USER"])
		return nil, err
	}

	logger.NewInfoLog(logger.JOURNEY["FIND_SERVICE"], logger.MESSAGE["OK"]["FOUND"])
	return user, nil
}

func (us *tUserService) findUserByEmailAndPassword(credentials domain.IUser) *ierrors.TError {
	logger.NewInfoLog(logger.JOURNEY["FIND_SERVICE"], logger.MESSAGE["INIT"]["FIND_EMAIL_PASS"])

	err := us.repository.FindUserByEmailAndPassword(credentials)

	if err != nil {
		logger.NewErrorLog(logger.JOURNEY["FIND_SERVICE"], logger.MESSAGE["ERROR"]["NO_USER"])
	}

	logger.NewInfoLog(logger.JOURNEY["FIND_SERVICE"], logger.MESSAGE["OK"]["FOUND"])
	return err
}
