package routes

import (
	"github.com/gin-gonic/gin"
	uc "github.com/william-cesar/crud-in-go/src/controller"
	m "github.com/william-cesar/crud-in-go/src/controller/middlewares"
)

func InitUserRoutes(r *gin.RouterGroup, uc uc.IUserController) {
	base := baseUrl + "/user"
	userById := base + "/:id"

	r.Use(m.Auth)

	r.POST(base, uc.FindUserByEmail)
	r.GET(userById, uc.FindUserById)
	r.PATCH(userById, uc.UpdateUser)
	r.DELETE(userById, uc.DeleteUser)
}
