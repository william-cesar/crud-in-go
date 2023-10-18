package domain

import (
	"crypto/sha256"
	"encoding/hex"
)

func (u *tUser) EncryptPassword() {
	hash := sha256.New()
	defer hash.Reset()

	hash.Write([]byte(u.password))
	u.password = hex.EncodeToString(hash.Sum(nil))
}
