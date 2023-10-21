package handlers

import (
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/config/logger"
)

func RemoveBrearePrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		return token[7:]
	}

	return token
}

func VerifyToken(tk string) *ierrors.TError {
	logger.NewInfoLog(logger.JOURNEY["AUTH"], logger.MESSAGE["INIT"]["TOKEN_VERIFY"])

	secret := os.Getenv("JWT_SEED")
	trimmedToken := RemoveBrearePrefix(tk)

	token, err := jwt.Parse(trimmedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.NewErrorLog(logger.JOURNEY["AUTH"], logger.MESSAGE["ERROR"]["TOKEN_SIGNATURE"])
			return nil, ierrors.NewBadRequestError("Invalid token")
		}
		return []byte(secret), nil
	})

	if err != nil {
		logger.NewErrorLog(logger.JOURNEY["AUTH"], logger.MESSAGE["ERROR"]["TOKEN_VERIFY"])
		return ierrors.NewUnauthorizedError()
	}

	_, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		logger.NewErrorLog(logger.JOURNEY["AUTH"], logger.MESSAGE["ERROR"]["TOKEN_VALIDITY"])
		return ierrors.NewUnauthorizedError()
	}

	logger.NewInfoLog(logger.JOURNEY["AUTH"], logger.MESSAGE["OK"]["TOKEN_VERIFY"])
	return nil
}
