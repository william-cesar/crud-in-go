package repository

import (
	"context"

	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/config/logger"
	"github.com/william-cesar/crud-in-go/src/model/adapters"
	"github.com/william-cesar/crud-in-go/src/model/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *tUserRepository) CreateUser(d domain.IUser) (domain.IUser, *ierrors.TError) {
	logger.NewInfoLog(logger.JOURNEY["CREATE_REPOSITORY"], logger.MESSAGE["INIT"]["CREATION"])

	collection := ur.dbconn.Collection(COLLECTION)

	dbEntity := adapters.ConvertDomainToEntity(d)

	res, err := collection.InsertOne(context.Background(), dbEntity)

	if err != nil {
		logger.NewErrorLog(logger.JOURNEY["CREATE_REPOSITORY"], logger.MESSAGE["ERROR"]["CREATION"], err)
		return nil, ierrors.NewInternalError(err.Error())
	}

	dbEntity.ID = res.InsertedID.(primitive.ObjectID)

	logger.NewInfoLog(logger.JOURNEY["CREATE_REPOSITORY"], logger.MESSAGE["OK"]["CREATED"])
	return adapters.ConvertEntityToDomain(*dbEntity), nil
}
