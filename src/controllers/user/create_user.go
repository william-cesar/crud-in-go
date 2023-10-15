package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/william-cesar/crud-in-go/src/config/validations"
	"github.com/william-cesar/crud-in-go/src/view/adapters"
	"github.com/william-cesar/crud-in-go/src/view/domain"
)

func (uc *tUserController) CreateUser(c *gin.Context) {
	var userRequest adapters.TCreateUserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		reqErr := validations.ValidateUserError(err)
		c.JSON(reqErr.StatusCode, reqErr)
		return
	}

	newUser := domain.NewUser(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	uc.service.CreateUserService(newUser)

	c.JSON(http.StatusCreated, nil)
}
