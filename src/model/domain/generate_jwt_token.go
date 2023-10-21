package domain

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
)

func (u *tUser) GenerateToken() (string, *ierrors.TError) {
	secret := os.Getenv("JWT_SEED")

	claims := jwt.MapClaims{
		"email": u.encryptEmail(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", ierrors.NewInternalError()
	}

	return tokenString, nil
}
