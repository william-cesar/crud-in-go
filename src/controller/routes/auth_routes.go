package routes

import "github.com/gin-gonic/gin"

func InitAuthRoutes(url string, r *gin.RouterGroup) {
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
