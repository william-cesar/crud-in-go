package repository

import (
	"context"

	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	ma "github.com/william-cesar/crud-in-go/src/model/adapters"
	d "github.com/william-cesar/crud-in-go/src/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *tUserRepository) UpdateUser(id string, user d.IUser) *ierrors.TError {
	collection := ur.dbconn.Collection(COLLECTION)

	dbEntity := ma.ConvertDomainToEntity(user)

	userId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: userId}}
	fields := bson.D{{Key: "$set", Value: dbEntity}}

	if _, err := collection.UpdateOne(context.Background(), filter, fields); err != nil {
		return ierrors.NewInternalError()
	}

	return nil
}
