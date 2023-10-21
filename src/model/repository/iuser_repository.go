package repository

import (
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/model/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION = "users"

func NewUserRepository(dbconn *mongo.Database) IUserRepository {
	return &tUserRepository{
		dbconn: dbconn,
	}
}

type tUserRepository struct {
	dbconn *mongo.Database
}

type IUserRepository interface {
	CreateUser(domain.IUser) (domain.IUser, *ierrors.TError)
	FindUserByEmail(email string) (domain.IUser, *ierrors.TError)
	FindUserById(email string) (domain.IUser, *ierrors.TError)
	DeleteUser(id string) *ierrors.TError
	ActivateUser(id string) *ierrors.TError
	UpdateUser(id string, user domain.IUser) *ierrors.TError

	FindUserByEmailAndPassword(credentials domain.IUser) *ierrors.TError
}
