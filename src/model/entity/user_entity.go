package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type TuserEntity struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Age      int8               `bson:"age,omitempty"`
	Active   bool               `bson:"active,omitempty"`
}
