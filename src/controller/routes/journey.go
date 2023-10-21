package routes

import (
	"github.com/gin-gonic/gin"
	uc "github.com/william-cesar/crud-in-go/src/controller"
)

func InitAuthRoutes(r *gin.RouterGroup, uc uc.IUserController) {
	signup := baseUrl + "/signup"
	login := baseUrl + "/login"
	activate := baseUrl + "/activate/:id"

	r.POST(login, uc.Login)
	r.POST(signup, uc.CreateUser)
	r.GET(activate, uc.ActivateUser)
}
