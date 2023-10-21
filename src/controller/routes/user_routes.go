package routes

import (
	"github.com/gin-gonic/gin"
	uc "github.com/william-cesar/crud-in-go/src/controller"
)

func InitUserRoutes(r *gin.RouterGroup, uc uc.IUserController) {
	base := baseUrl + "/users"
	user := base + "/user"
	userById := base + "/user/:id"

	r.GET(userById, uc.FindUserById)
	r.POST(user, uc.FindUserByEmail)
	r.POST(base, uc.CreateUser)
	r.PATCH(userById, uc.UpdateUser)
	r.DELETE(userById, uc.DeleteUser)
}
