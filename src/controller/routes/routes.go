package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	uc "github.com/william-cesar/crud-in-go/src/controller"
)

var baseUrl string

func InitRoutes(r *gin.RouterGroup, uc uc.IUserController) {
	baseUrl = fmt.Sprintf("/%s", os.Getenv("ENTRYPOINT"))

	InitAuthRoutes(r, uc)
	InitUserRoutes(r, uc)
}
