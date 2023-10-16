package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type TuserEntity struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Name     string             `bson:"name"`
	Age      int8               `bson:"age"`
	Active   bool               `bson:"active"`
}
