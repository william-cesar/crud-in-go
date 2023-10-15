package routes

import (
	"github.com/gin-gonic/gin"
	uc "github.com/william-cesar/crud-in-go/src/controller/user"
)

func InitUserRoutes(url string, r *gin.RouterGroup) {
	user := url + "/user"
	userById := user + "/:id"

	r.GET(userById, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "User found",
		})
	})

	r.POST(user, uc.CreateUser)
}
