package repository

import (
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	d "github.com/william-cesar/crud-in-go/src/model/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(dbconn *mongo.Database) IUserRepository {
	return &tUserRepository{
		dbconn: dbconn,
	}
}

type tUserRepository struct {
	dbconn *mongo.Database
}

type IUserRepository interface {
	CreateUser(d.IUser) (d.IUser, *ierrors.TError)
}
