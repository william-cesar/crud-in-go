package repository

import (
	"context"

	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/config/logger"
	"github.com/william-cesar/crud-in-go/src/model/adapters"
	"github.com/william-cesar/crud-in-go/src/model/domain"
	"github.com/william-cesar/crud-in-go/src/model/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const activateWarning = "User is not activated. Please verify your email."

func (ur *tUserRepository) FindUserByEmail(email string) (domain.IUser, *ierrors.TError) {
	logger.NewInfoLog(logger.JOURNEY["FIND_REPOSITORY"], logger.MESSAGE["INIT"]["FIND_EMAIL"])

	collection := ur.dbconn.Collection(COLLECTION)

	dbEntity := &entity.TuserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(dbEntity)

	if err != nil {
		logger.NewErrorLog(logger.JOURNEY["FIND_REPOSITORY"], logger.MESSAGE["ERROR"]["NO_USER"])

		if err == mongo.ErrNoDocuments {
			return nil, ierrors.NewNotFoundError()
		}

		return nil, ierrors.NewInternalError()
	}

	if !dbEntity.Active {
		logger.NewErrorLog(logger.JOURNEY["FIND_REPOSITORY"], logger.MESSAGE["ERROR"]["ACTIVE"])
		return nil, ierrors.NewBadRequestError(activateWarning)
	}

	logger.NewInfoLog(logger.JOURNEY["FIND_REPOSITORY"], logger.MESSAGE["OK"]["FOUND"])
	return adapters.ConvertEntityToDomain(*dbEntity), nil
}

func (ur *tUserRepository) FindUserById(id string) (domain.IUser, *ierrors.TError) {
	logger.NewInfoLog(logger.JOURNEY["FIND_REPOSITORY"], logger.MESSAGE["INIT"]["FIND_ID"])

	collection := ur.dbconn.Collection(COLLECTION)

	dbEntity := &entity.TuserEntity{}

	userId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		logger.NewErrorLog(logger.JOURNEY["FIND_REPOSITORY"], logger.MESSAGE["ERROR"]["INVALID_ID"])
		return nil, ierrors.NewBadRequestError()
	}

	filter := bson.D{{Key: "_id", Value: userId}}
	err = collection.FindOne(context.Background(), filter).Decode(dbEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.NewWarningLog(logger.JOURNEY["FIND_REPOSITORY"], logger.MESSAGE["ERROR"]["NO_USER"])
			return nil, ierrors.NewNotFoundError()
		}

		logger.NewErrorLog(logger.JOURNEY["FIND_REPOSITORY"], logger.MESSAGE["ERROR"]["NO_USER"])
		return nil, ierrors.NewInternalError()
	}

	if !dbEntity.Active {
		logger.NewWarningLog(logger.JOURNEY["FIND_REPOSITORY"], logger.MESSAGE["ERROR"]["ACTIVE"])
		return nil, ierrors.NewBadRequestError(activateWarning)
	}

	logger.NewInfoLog(logger.JOURNEY["FIND_REPOSITORY"], logger.MESSAGE["OK"]["FOUND"])
	return adapters.ConvertEntityToDomain(*dbEntity), nil
}

func (ur *tUserRepository) FindUserByEmailAndPassword(credentials domain.IUser) *ierrors.TError {
	logger.NewInfoLog(logger.JOURNEY["FIND_REPOSITORY"], logger.MESSAGE["INIT"]["FIND_EMAIL_PASS"])

	collection := ur.dbconn.Collection(COLLECTION)

	dbEntity := &entity.TuserEntity{}

	filter := bson.D{
		{Key: "email", Value: credentials.GetEmail()},
		{Key: "password", Value: credentials.GetPassword()},
	}
	err := collection.FindOne(context.Background(), filter).Decode(dbEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.NewWarningLog(logger.JOURNEY["FIND_REPOSITORY"], logger.MESSAGE["ERROR"]["NO_USER"])
			return ierrors.NewBadRequestError("Invalid credentials.")
		}

		logger.NewErrorLog(logger.JOURNEY["FIND_REPOSITORY"], logger.MESSAGE["ERROR"]["NO_USER"])
		return ierrors.NewInternalError()
	}

	if !dbEntity.Active {
		logger.NewWarningLog(logger.JOURNEY["FIND_REPOSITORY"], logger.MESSAGE["ERROR"]["ACTIVE"])
		return ierrors.NewBadRequestError(activateWarning)
	}

	logger.NewInfoLog(logger.JOURNEY["FIND_REPOSITORY"], logger.MESSAGE["OK"]["FOUND"])
	return nil
}
