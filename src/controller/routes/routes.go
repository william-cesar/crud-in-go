package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/william-cesar/crud-in-go/docs"
	uc "github.com/william-cesar/crud-in-go/src/controller"
)

var baseUrl string

func InitRoutes(r *gin.RouterGroup, uc uc.IUserController) {
	baseUrl = fmt.Sprintf("/%s", os.Getenv("ENTRYPOINT"))

	// Swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	InitAuthRoutes(r, uc)
	InitUserRoutes(r, uc)
}
