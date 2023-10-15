package routes

import (
	"github.com/gin-gonic/gin"
	uc "github.com/william-cesar/crud-in-go/src/controllers/user"
	us "github.com/william-cesar/crud-in-go/src/model/services/user"
)

func InitUserRoutes(url string, r *gin.RouterGroup) {
	user := url + "/user"
	userById := user + "/:id"

	userController := uc.NewUserController(us.NewUserService())

	r.GET(userById, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "User found",
		})
	})

	r.POST(user, userController.CreateUser)
}
