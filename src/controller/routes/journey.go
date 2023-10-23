package routes

import (
	"github.com/gin-gonic/gin"
	uc "github.com/william-cesar/crud-in-go/src/controller"
)

func InitAuthRoutes(r *gin.RouterGroup, uc uc.IUserController) {
	r.POST("signup", uc.CreateUser)
	r.POST("login", uc.Login)
	r.GET("activate/:id", uc.ActivateUser)
}
