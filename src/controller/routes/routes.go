package routes

import "github.com/gin-gonic/gin"

const (
	baseUrl  = "/api/v1"
	user     = baseUrl + "/user"
	userById = user + "/:id"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET(userById, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
}
