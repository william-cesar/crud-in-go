package routes

import (
	"github.com/gin-gonic/gin"
	uc "github.com/william-cesar/crud-in-go/src/controller"
)

func InitUserRoutes(r *gin.RouterGroup, uc uc.IUserController) {
	base := baseUrl + "/user"
	userById := base + "/:id"

	r.POST(base, uc.FindUserByEmail)
	r.GET(userById, uc.FindUserById)
	r.PATCH(userById, uc.UpdateUser)
	r.DELETE(userById, uc.DeleteUser)
}
