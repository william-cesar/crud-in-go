package service

import (
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/config/logger"
	"github.com/william-cesar/crud-in-go/src/model/domain"
)

func (us *tUserService) LoginUserService(user domain.IUser) (string, *ierrors.TError) {
	logger.NewInfoLog(logger.JOURNEY["LOGIN_SERVICE"], logger.MESSAGE["INIT"]["LOGIN"])

	user.EncryptPassword()

	if err := us.findUserByEmailAndPassword(user); err != nil {
		logger.NewErrorLog(logger.JOURNEY["LOGIN_SERVICE"], logger.MESSAGE["ERROR"]["NO_USER"], err)
		return "", err
	}

	logger.NewInfoLog(logger.JOURNEY["LOGIN_SERVICE"], logger.MESSAGE["OK"]["LOGGED"])
	return user.GenerateToken()
}
