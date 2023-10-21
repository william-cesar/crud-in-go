package repository

import (
	"context"

	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/model/adapters"
	"github.com/william-cesar/crud-in-go/src/model/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *tUserRepository) CreateUser(d domain.IUser) (domain.IUser, *ierrors.TError) {
	collection := ur.dbconn.Collection(COLLECTION)

	dbEntity := adapters.ConvertDomainToEntity(d)

	res, err := collection.InsertOne(context.Background(), dbEntity)

	if err != nil {
		return nil, ierrors.NewInternalError(err.Error())
	}

	dbEntity.ID = res.InsertedID.(primitive.ObjectID)

	return adapters.ConvertEntityToDomain(*dbEntity), nil
}
