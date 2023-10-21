package service

import (
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/config/logger"
)

func (us *tUserService) ActivateUserService(id string) *ierrors.TError {
	logger.NewInfoLog(logger.JOURNEY["ACTIVATE_SERVICE"], logger.MESSAGE["INIT"]["ACTIVATE"])

	err := us.repository.ActivateUser(id)

	if err != nil {
		logger.NewErrorLog(logger.JOURNEY["ACTIVATE_SERVICE"], logger.MESSAGE["ERROR"]["ACTIVATION"])
	}

	logger.NewInfoLog(logger.JOURNEY["ACTIVATE_SERVICE"], logger.MESSAGE["OK"]["ACTIVE"])
	return err
}
