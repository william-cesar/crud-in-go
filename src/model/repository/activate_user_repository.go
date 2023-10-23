package repository

import (
	"context"

	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/config/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *tUserRepository) ActivateUser(id string) *ierrors.TError {
	logger.NewInfoLog(logger.JOURNEY["ACTIVATE_REPOSITORY"], logger.MESSAGE["INIT"]["ACTIVATION"])

	collection := ur.dbconn.Collection(COLLECTION)

	userId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		logger.NewErrorLog(logger.JOURNEY["ACTIVATE_REPOSITORY"], logger.MESSAGE["ERROR"]["INVALID_ID"], err)
		return ierrors.NewBadRequestError()
	}

	filter := bson.D{{Key: "_id", Value: userId}}
	field := bson.D{{Key: "$set", Value: bson.D{{Key: "active", Value: true}}}}

	res, err := collection.UpdateOne(context.Background(), filter, field)

	if res.ModifiedCount == 0 {
		if res.MatchedCount == 1 {
			logger.NewWarningLog(logger.JOURNEY["ACTIVATE_REPOSITORY"], logger.MESSAGE["WARN"]["ACTIVE"])
			return ierrors.NewBadRequestError("User already activated.")
		}

		logger.NewErrorLog(logger.JOURNEY["ACTIVATE_REPOSITORY"], logger.MESSAGE["ERROR"]["NO_USER"], nil)
		return ierrors.NewNotFoundError()
	}

	if err != nil {
		logger.NewErrorLog(logger.JOURNEY["ACTIVATE_REPOSITORY"], logger.MESSAGE["ERROR"]["ACTIVATION"], err)
		return ierrors.NewInternalError()
	}

	logger.NewInfoLog(logger.JOURNEY["ACTIVATE_REPOSITORY"], logger.MESSAGE["OK"]["ACTIVE"])
	return nil
}
