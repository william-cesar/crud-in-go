package repository

import (
	"context"

	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *tUserRepository) DeleteUser(id string) *ierrors.TError {
	collection := ur.dbconn.Collection(COLLECTION)

	userId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: userId}}

	res, err := collection.DeleteOne(context.Background(), filter)

	if res.DeletedCount == 0 {
		return ierrors.NewNotFoundError()
	}

	if err != nil {
		return ierrors.NewInternalError()
	}

	return nil
}
