package service

import (
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/config/logger"
)

func (us *tUserService) DeleteUserService(id string) *ierrors.TError {
	logger.NewInfoLog(logger.JOURNEY["DELETE_SERVICE"], logger.MESSAGE["INIT"]["DELETION"])

	err := us.repository.DeleteUser(id)

	if err != nil {
		logger.NewErrorLog(logger.JOURNEY["DELETE_SERVICE"], logger.MESSAGE["ERROR"]["DELETION"])
	}

	logger.NewInfoLog(logger.JOURNEY["DELETE_SERVICE"], logger.MESSAGE["OK"]["DELETED"])
	return err
}
