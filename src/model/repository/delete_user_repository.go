package repository

import (
	"context"

	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/config/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *tUserRepository) DeleteUser(id string) *ierrors.TError {
	logger.NewInfoLog(logger.JOURNEY["DELETE_REPOSITORY"], logger.MESSAGE["INIT"]["DELETION"])

	collection := ur.dbconn.Collection(COLLECTION)

	userId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		logger.NewErrorLog(logger.JOURNEY["DELETE_REPOSITORY"], logger.MESSAGE["ERROR"]["INVALID_ID"])
		return ierrors.NewBadRequestError()
	}

	filter := bson.D{{Key: "_id", Value: userId}}

	res, err := collection.DeleteOne(context.Background(), filter)

	if res.DeletedCount == 0 {
		logger.NewErrorLog(logger.JOURNEY["DELETE_REPOSITORY"], logger.MESSAGE["ERROR"]["NO_USER"])
		return ierrors.NewNotFoundError()
	}

	if err != nil {
		logger.NewErrorLog(logger.JOURNEY["DELETE_REPOSITORY"], logger.MESSAGE["ERROR"]["DELETION"])
		return ierrors.NewInternalError()
	}

	logger.NewInfoLog(logger.JOURNEY["DELETE_REPOSITORY"], logger.MESSAGE["OK"]["DELETED"])
	return nil
}
