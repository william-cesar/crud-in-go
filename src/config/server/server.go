package server

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/william-cesar/crud-in-go/src/config/database/mongodb"
	"github.com/william-cesar/crud-in-go/src/config/logger"
	"github.com/william-cesar/crud-in-go/src/controller/routes"
)

func Init() {
	logger.NewInfoLog(logger.JOURNEY["SERVER"], logger.MESSAGE["INIT"]["SERVER"])

	if err := godotenv.Load(); err != nil {
		logger.NewErrorLog(logger.JOURNEY["SERVER"], logger.MESSAGE["ERROR"]["LOAD_ENV"], err)
		os.Exit(1)
	}

	url := os.Getenv("PORT")

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.SetTrustedProxies([]string{"*"})

	dbconn, err := mongodb.NewMongoDBConnection(context.Background())

	if err != nil {
		logger.NewErrorLog(logger.JOURNEY["SERVER"], logger.MESSAGE["ERROR"]["DB_CONN"], err)
		os.Exit(1)
	}

	uc := InitDependencies(dbconn)
	InitDocsDependencies()

	routes.InitRoutes(router, uc)

	if err := router.Run(url); err != nil {
		logger.NewErrorLog(logger.JOURNEY["SERVER"], logger.MESSAGE["ERROR"]["SERVER"], err)
		os.Exit(1)
	}
}
