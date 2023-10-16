package server

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/william-cesar/crud-in-go/src/config/database/mongodb"
	"github.com/william-cesar/crud-in-go/src/controller/routes"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	url := os.Getenv("URL")

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.SetTrustedProxies([]string{"*"})

	dbconn, err := mongodb.NewMongoDBConnection(context.Background())

	if err != nil {
		log.Fatal(err)
		return
	}

	uc := InitDependencies(dbconn)

	routes.InitRoutes(&router.RouterGroup, uc)

	if err := router.Run(url); err != nil {
		log.Fatal(err)
	}
}
