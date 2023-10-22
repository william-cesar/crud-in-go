package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/william-cesar/crud-in-go/src/config/logger"
	"github.com/william-cesar/crud-in-go/src/config/validations"
	"github.com/william-cesar/crud-in-go/src/model/domain"
	"github.com/william-cesar/crud-in-go/src/view/adapters"
)

// Login  				godoc
// @Summary 			Log users in
// @Description 	Login and set token in header
// @Param 				json body adapters.TUserLoginRequest true "user"
// @Accept 				application/json
// @Produce 			application/json
// @Tags 					journey
// @Success 			200 {object} adapters.TSuccessResponse
// @Failure 			400 {object} ierrors.TError
// @Router 				/login  [post]
func (uc *tUserController) Login(c *gin.Context) {
	logger.NewInfoLog(logger.JOURNEY["LOGIN_CONTROLLER"], logger.MESSAGE["INIT"]["LOGIN"])

	var userLogin adapters.TUserLoginRequest

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		logger.NewErrorLog(logger.JOURNEY["LOGIN_CONTROLLER"], logger.MESSAGE["ERROR"]["REQ_BODY"])
		reqErr := validations.ValidateUserError(err)
		c.JSON(reqErr.StatusCode, reqErr)
		return
	}

	credentials := domain.NewUserLogin(
		userLogin.Email,
		userLogin.Password,
	)

	token, err := uc.service.LoginUserService(credentials)

	if err != nil {
		logger.NewErrorLog(logger.JOURNEY["LOGIN_CONTROLLER"], logger.MESSAGE["ERROR"]["TOKEN_VERIFY"])
		c.JSON(err.StatusCode, err)
		return
	}

	c.Header("Authorization", token)

	res := adapters.TSuccessResponse{
		Message: "User logged in successfully.",
	}

	logger.NewInfoLog(logger.JOURNEY["LOGIN_CONTROLLER"], logger.MESSAGE["OK"]["LOGGED"])
	c.JSON(http.StatusCreated, res)
}
