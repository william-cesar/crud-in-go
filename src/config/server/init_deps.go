package server

import (
	uc "github.com/william-cesar/crud-in-go/src/controller"
	ur "github.com/william-cesar/crud-in-go/src/model/repository"
	us "github.com/william-cesar/crud-in-go/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitDependencies(db *mongo.Database) uc.IUserController {
	repo := ur.NewUserRepository(db)
	service := us.NewUserService(repo)

	return uc.NewUserController(service)
}
