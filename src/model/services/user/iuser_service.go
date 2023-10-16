package user

import (
	d "github.com/william-cesar/crud-in-go/src/view/domain"
)

func NewUserService() IUserService {
	return &tUserService{}
}

type tUserService struct{}

type IUserService interface {
	CreateUserService(d.IUser)
	// GetUserById()
	// UpdateUser()
	// DeleteUser()
}
