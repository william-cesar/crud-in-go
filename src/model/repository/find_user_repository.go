package repository

import (
	"context"
	"fmt"

	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	ma "github.com/william-cesar/crud-in-go/src/model/adapters"
	d "github.com/william-cesar/crud-in-go/src/model/domain"
	"github.com/william-cesar/crud-in-go/src/model/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur *tUserRepository) FindUser(email string) (d.IUser, *ierrors.TError) {
	collection := ur.dbconn.Collection(COLLECTION)

	dbEntity := &entity.TuserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(dbEntity)

	if err != nil {
		fmt.Println(err)
		if err == mongo.ErrNoDocuments {
			return nil, ierrors.NewNotFoundError()
		}
		return nil, ierrors.NewInternalError()
	}

	return ma.ConvertEntityToDomain(*dbEntity), nil

}
