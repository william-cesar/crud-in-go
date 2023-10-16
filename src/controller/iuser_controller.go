package controller

import (
	"github.com/gin-gonic/gin"
	us "github.com/william-cesar/crud-in-go/src/model/service"
)

func NewUserController(iservice us.IUserService) IUserController {
	return &tUserController{
		service: iservice,
	}
}

type IUserController interface {
	CreateUser(c *gin.Context)
	// GetUserById(c *gin.Context)
	// UpdateUser(c *gin.Context)
	// DeleteUser(c *gin.Context)
}

type tUserController struct {
	service us.IUserService
}
