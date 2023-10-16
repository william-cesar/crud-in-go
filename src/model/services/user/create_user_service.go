package user

import d "github.com/william-cesar/crud-in-go/src/model/domain"

func (us *tUserService) CreateUserService(newUser d.IUser) {
	newUser.EncryptPassword()
}
