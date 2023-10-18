package routes

import (
	"github.com/gin-gonic/gin"
	uc "github.com/william-cesar/crud-in-go/src/controller"
)

func InitUserRoutes(url string, r *gin.RouterGroup, uc uc.IUserController) {
	base := url + "/users"
	user := base + "/user"
	userById := base + "/user/:id"

	r.GET(userById, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "User found",
		})
	})

	r.POST(user, uc.FindUser)
	r.POST(base, uc.CreateUser)
	r.DELETE(userById, uc.DeleteUser)
}
