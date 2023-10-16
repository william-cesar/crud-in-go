package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/william-cesar/crud-in-go/src/config/database/mongodb"
	"github.com/william-cesar/crud-in-go/src/controllers/routes"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	url := os.Getenv("URL")

	if url == "" {
		log.Fatal("URL is not set")
		return
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.SetTrustedProxies([]string{"*"})

	routes.InitRoutes(&router.RouterGroup)

	mongodb.InitConnection()

	if err := router.Run(url); err != nil {
		log.Fatal(err)
	}
}
