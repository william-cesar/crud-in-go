package service

import (
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	d "github.com/william-cesar/crud-in-go/src/model/domain"
	ur "github.com/william-cesar/crud-in-go/src/model/repository"
)

func NewUserService(irepository ur.IUserRepository) IUserService {
	return &tUserService{
		repository: irepository,
	}
}

type tUserService struct {
	repository ur.IUserRepository
}

type IUserService interface {
	CreateUserService(d.IUser) (d.IUser, *ierrors.TError)
	FindUserService(email string) (d.IUser, *ierrors.TError)
	// GetUserById()
	// UpdateUser()
	// DeleteUser()
}
