package repository

import (
	"context"

	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *tUserRepository) ActivateUser(id string) *ierrors.TError {
	collection := ur.dbconn.Collection(COLLECTION)

	userId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: userId}}
	field := bson.D{{Key: "$set", Value: bson.D{{Key: "active", Value: true}}}}

	res, err := collection.UpdateOne(context.Background(), filter, field)

	if res.ModifiedCount == 0 {
		if res.MatchedCount == 1 {
			return ierrors.NewBadRequestError("User already activated.")
		}

		return ierrors.NewNotFoundError()
	}

	if err != nil {
		return ierrors.NewInternalError()
	}

	return nil
}
