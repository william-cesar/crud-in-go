package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/william-cesar/crud-in-go/src/config/logger"
	"github.com/william-cesar/crud-in-go/src/config/validations"
	"github.com/william-cesar/crud-in-go/src/model/domain"
	"github.com/william-cesar/crud-in-go/src/view/adapters"
)

// Signup  				godoc
// @Summary 			Create users
// @Description 	Create user
// @Param 				json body adapters.TUserRequest true "user"
// @Accept 				application/json
// @Produce 			application/json
// @Tags 					journey
// @Success 			201 {object} adapters.TUserResponse
// @Failure 			400 {object} ierrors.TError
// @Router 				/signup  [post]
func (uc *tUserController) CreateUser(c *gin.Context) {
	logger.NewInfoLog(logger.JOURNEY["CREATE_CONTROLLER"], logger.MESSAGE["INIT"]["CREATION"])

	var userRequest adapters.TUserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.NewErrorLog(logger.JOURNEY["CREATE_CONTROLLER"], logger.MESSAGE["ERROR"]["INVALID_REQ"], err)
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
		logger.NewErrorLog(logger.JOURNEY["CREATE_CONTROLLER"], logger.MESSAGE["ERROR"]["CREATION"], err)
		c.JSON(err.StatusCode, err)
		return
	}

	logger.NewInfoLog(logger.JOURNEY["CREATE_CONTROLLER"], logger.MESSAGE["OK"]["CREATED"])
	c.JSON(http.StatusCreated, adapters.ConvertDomainToResponse(res))
}
