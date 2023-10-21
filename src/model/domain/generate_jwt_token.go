package domain

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/config/logger"
)

func (u *tUser) GenerateToken() (string, *ierrors.TError) {
	logger.NewInfoLog(logger.JOURNEY["AUTH"], logger.MESSAGE["INIT"]["TOKEN_GENERATION"])

	secret := os.Getenv("JWT_SEED")

	claims := jwt.MapClaims{
		"email": u.encryptEmail(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		logger.NewErrorLog(logger.JOURNEY["AUTH"], logger.MESSAGE["ERROR"]["TOKEN_GENERATION"])
		return "", ierrors.NewInternalError()
	}

	logger.NewInfoLog(logger.JOURNEY["AUTH"], logger.MESSAGE["OK"]["TOKEN_GENERATION"])
	return tokenString, nil
}
