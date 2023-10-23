package repository

import (
	"context"

	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/config/logger"
	"github.com/william-cesar/crud-in-go/src/model/adapters"
	"github.com/william-cesar/crud-in-go/src/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *tUserRepository) UpdateUser(id string, user domain.IUser) *ierrors.TError {
	logger.NewInfoLog(logger.JOURNEY["UPDATE_REPOSITORY"], logger.MESSAGE["INIT"]["UPDATE"])

	collection := ur.dbconn.Collection(COLLECTION)

	dbEntity := adapters.ConvertDomainToEntity(user)

	userId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		logger.NewErrorLog(logger.JOURNEY["UPDATE_REPOSITORY"], logger.MESSAGE["ERROR"]["INVALID_ID"], err)
		return ierrors.NewBadRequestError()
	}

	filter := bson.D{{Key: "_id", Value: userId}}
	fields := bson.D{{Key: "$set", Value: dbEntity}}

	if _, err := collection.UpdateOne(context.Background(), filter, fields); err != nil {
		logger.NewErrorLog(logger.JOURNEY["UPDATE_REPOSITORY"], logger.MESSAGE["ERROR"]["UPDATE"], err)
		return ierrors.NewInternalError()
	}

	logger.NewInfoLog(logger.JOURNEY["UPDATE_REPOSITORY"], logger.MESSAGE["OK"]["UPDATED"])
	return nil
}
