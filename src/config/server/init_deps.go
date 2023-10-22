package server

import (
	"os"

	docs "github.com/william-cesar/crud-in-go/docs"
	uc "github.com/william-cesar/crud-in-go/src/controller"
	ur "github.com/william-cesar/crud-in-go/src/model/repository"
	us "github.com/william-cesar/crud-in-go/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitDocsDependencies() {
	host, version := os.Getenv("BASE_URL"), os.Getenv("ENTRYPOINT")
	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.BasePath = "/" + version
	docs.SwaggerInfo.Title = "Users API"
	docs.SwaggerInfo.Description = "An API developed in GoLang using Gin framework which allows users to perform CRUD operations."
	docs.SwaggerInfo.Version = "1.0"
}

func InitDependencies(db *mongo.Database) uc.IUserController {
	repo := ur.NewUserRepository(db)
	service := us.NewUserService(repo)

	return uc.NewUserController(service)
}
