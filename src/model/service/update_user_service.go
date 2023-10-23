package service

import (
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/config/logger"
	"github.com/william-cesar/crud-in-go/src/model/domain"
)

func (us *tUserService) UpdateUserService(id string, user domain.IUser) *ierrors.TError {
	logger.NewInfoLog(logger.JOURNEY["UPDATE_SERVICE"], logger.MESSAGE["INIT"]["UPDATE"])

	dbUser, err := us.repository.FindUserById(id)

	if err != nil {
		logger.NewErrorLog(logger.JOURNEY["UPDATE_SERVICE"], logger.MESSAGE["ERROR"]["NO_USER"], err)
		return err
	}

	if dbUser != nil && !dbUser.GetActive() {
		logger.NewErrorLog(logger.JOURNEY["UPDATE_SERVICE"], logger.MESSAGE["ERROR"]["ACTIVE"], nil)
		return ierrors.NewBadRequestError("User is not activated. Please verify your email.")
	}

	if user.GetPassword() != "" {
		user.EncryptPassword()
	}

	if err := us.repository.UpdateUser(id, user); err != nil {
		logger.NewErrorLog(logger.JOURNEY["UPDATE_SERVICE"], logger.MESSAGE["ERROR"]["UPDATE"], err)
		return err
	}

	logger.NewInfoLog(logger.JOURNEY["UPDATE_SERVICE"], logger.MESSAGE["OK"]["UPDATED"])
	return nil
}
