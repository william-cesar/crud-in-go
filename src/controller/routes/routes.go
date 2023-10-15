package routes

import "github.com/gin-gonic/gin"

const baseUrl = "/api/v1"

func InitRoutes(r *gin.RouterGroup) {
	InitAuthRoutes(baseUrl, r)
	InitUserRoutes(baseUrl, r)
}
