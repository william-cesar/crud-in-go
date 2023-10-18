package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/william-cesar/crud-in-go/src/config/validations"
	"github.com/william-cesar/crud-in-go/src/model/domain"
	"github.com/william-cesar/crud-in-go/src/view/adapters"
)

func (uc *tUserController) CreateUser(c *gin.Context) {
	var userRequest adapters.TUserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		reqErr := validations.ValidateUserError(err)
		c.JSON(reqErr.StatusCode, reqErr)
		return
	}

	defaultActive := false
	newUser := domain.NewUser(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
		defaultActive,
	)

	res, err := uc.service.CreateUserService(newUser)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusCreated, adapters.ConvertDomainToResponse(res))
}
