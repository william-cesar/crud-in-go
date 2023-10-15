package routes

import "github.com/gin-gonic/gin"

func InitUserRoutes(url string, r *gin.RouterGroup) {
	user := url + "/user"
	userById := user + "/:id"

	r.GET(userById, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "User found",
		})
	})
}
