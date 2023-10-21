package service

import (
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/model/domain"
	"github.com/william-cesar/crud-in-go/src/model/repository"
)

func NewUserService(irepository repository.IUserRepository) IUserService {
	return &tUserService{
		repository: irepository,
	}
}

type tUserService struct {
	repository repository.IUserRepository
}

type IUserService interface {
	CreateUserService(domain.IUser) (domain.IUser, *ierrors.TError)
	FindUserByEmailService(email string) (domain.IUser, *ierrors.TError)
	FindUserByIdService(email string) (domain.IUser, *ierrors.TError)
	DeleteUserService(id string) *ierrors.TError
	ActivateUserService(id string) *ierrors.TError
	UpdateUserService(id string, user domain.IUser) *ierrors.TError

	findUserByEmailAndPassword(credentials domain.IUser) *ierrors.TError

	LoginUserService(credentials domain.IUser) (string, *ierrors.TError)
}
