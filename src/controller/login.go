package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/william-cesar/crud-in-go/src/config/validations"
	"github.com/william-cesar/crud-in-go/src/model/domain"
	"github.com/william-cesar/crud-in-go/src/view/adapters"
)

func (uc *tUserController) Login(c *gin.Context) {
	var userLogin adapters.TUserLoginRequest

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		reqErr := validations.ValidateUserError(err)
		c.JSON(reqErr.StatusCode, reqErr)
		return
	}

	credentials := domain.NewUserLogin(
		userLogin.Email,
		userLogin.Password,
	)

	err := uc.service.LoginUserService(credentials)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	res := adapters.TSuccessResponse{
		Message: "User logged in successfully.",
	}

	c.JSON(http.StatusCreated, res)
}
