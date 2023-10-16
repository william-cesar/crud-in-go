package repository

import (
	"context"
	"os"

	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	ma "github.com/william-cesar/crud-in-go/src/model/adapters"
	d "github.com/william-cesar/crud-in-go/src/model/domain"
)

func (ur *tUserRepository) CreateUser(d d.IUser) (d.IUser, *ierrors.TError) {
	collectionName := os.Getenv("DB_COLLECTION")
	collection := ur.dbconn.Collection(collectionName)

	dbEntity := ma.ConvertDomainToEntity(d)

	res, err := collection.InsertOne(context.Background(), dbEntity)

	if err != nil {
		return nil, ierrors.NewInternalError(err.Error())
	}

	d.SetID(res.InsertedID.(string))

	return ma.ConvertEntityToDomain(*dbEntity), nil
}
