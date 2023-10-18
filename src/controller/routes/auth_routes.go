package routes

import (
	"github.com/gin-gonic/gin"
	uc "github.com/william-cesar/crud-in-go/src/controller"
)

func InitAuthRoutes(url string, r *gin.RouterGroup, uc uc.IUserController) {
	signup := url + "/signup"
	login := url + "/login"

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
}