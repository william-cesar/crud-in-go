package routes

import (
	"github.com/gin-gonic/gin"
	uc "github.com/william-cesar/crud-in-go/src/controller"
)

const baseUrl = "/api/v1"

func InitRoutes(r *gin.RouterGroup, uc uc.IUserController) {
	InitAuthRoutes(baseUrl, r, uc)
	InitUserRoutes(baseUrl, r, uc)
}
