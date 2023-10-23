package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	uc "github.com/william-cesar/crud-in-go/src/controller"
)

func InitRoutes(r *gin.Engine, uc uc.IUserController) {
	baseUrl := fmt.Sprintf("/%s", os.Getenv("ENTRYPOINT"))

	journeyGroup := r.Group(baseUrl)
	userGroup := r.Group(baseUrl + "/user")

	// Swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
	})

	InitAuthRoutes(journeyGroup, uc)
	InitUserRoutes(userGroup, uc)
}
