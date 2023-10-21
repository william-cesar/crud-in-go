package routes

import (
	"github.com/gin-gonic/gin"
	uc "github.com/william-cesar/crud-in-go/src/controller"
)

func InitAuthRoutes(r *gin.RouterGroup, uc uc.IUserController) {
	signup := baseUrl + "/signup"
	login := baseUrl + "/login"
	activate := baseUrl + "/activate/:id"

	r.POST(signup, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Signed up",
		})
	})

	r.POST(login, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Logged in",
		})
	})

	r.GET(activate, uc.ActivateUser)
}
