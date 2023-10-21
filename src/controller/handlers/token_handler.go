package handlers

import (
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
)

func RemoveBrearePrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		return token[7:]
	}

	return token
}

func VerifyToken(tk string) *ierrors.TError {
	secret := os.Getenv("JWT_SEED")
	trimmedToken := RemoveBrearePrefix(tk)

	token, err := jwt.Parse(trimmedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ierrors.NewBadRequestError("Invalid token")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return ierrors.NewUnauthorizedError()
	}

	_, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return ierrors.NewUnauthorizedError()
	}

	return nil
}
