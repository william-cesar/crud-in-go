package repository

import (
	"context"

	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	ma "github.com/william-cesar/crud-in-go/src/model/adapters"
	d "github.com/william-cesar/crud-in-go/src/model/domain"
	"github.com/william-cesar/crud-in-go/src/model/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur *tUserRepository) FindUserByEmail(email string) (d.IUser, *ierrors.TError) {
	collection := ur.dbconn.Collection(COLLECTION)

	dbEntity := &entity.TuserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(dbEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ierrors.NewNotFoundError()
		}
		return nil, ierrors.NewInternalError()
	}

	if !dbEntity.Active {
		return nil, ierrors.NewBadRequestError("User is not activated.")
	}

	return ma.ConvertEntityToDomain(*dbEntity), nil
}

func (ur *tUserRepository) FindUserById(id string) (d.IUser, *ierrors.TError) {
	collection := ur.dbconn.Collection(COLLECTION)

	dbEntity := &entity.TuserEntity{}

	userId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: userId}}
	err := collection.FindOne(context.Background(), filter).Decode(dbEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ierrors.NewNotFoundError()
		}
		return nil, ierrors.NewInternalError()
	}

	if !dbEntity.Active {
		return nil, ierrors.NewBadRequestError("User is not activated.")
	}

	return ma.ConvertEntityToDomain(*dbEntity), nil
}
