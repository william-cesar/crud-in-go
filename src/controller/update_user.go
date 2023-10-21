package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/config/logger"
	"github.com/william-cesar/crud-in-go/src/config/validations"
	"github.com/william-cesar/crud-in-go/src/model/domain"
	"github.com/william-cesar/crud-in-go/src/view/adapters"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *tUserController) UpdateUser(c *gin.Context) {
	logger.NewInfoLog(logger.JOURNEY["UPDATE_CONTROLLER"], logger.MESSAGE["INIT"]["UPDATE"])

	var uReq adapters.TUserUpdateRequest

	if err := c.ShouldBindJSON(&uReq); err != nil {
		logger.NewErrorLog(logger.JOURNEY["UPDATE_CONTROLLER"], logger.MESSAGE["ERROR"]["REQ_BODY"])
		reqErr := validations.ValidateUserError(err)
		c.JSON(reqErr.StatusCode, reqErr)
		return
	}

	id := c.Param("id")

	if _, err := primitive.ObjectIDFromHex(id); err != nil || strings.TrimSpace(id) == "" {
		logger.NewErrorLog(logger.JOURNEY["UPDATE_CONTROLLER"], logger.MESSAGE["ERROR"]["INVALID_ID"])
		e := ierrors.NewBadRequestError()
		c.JSON(e.StatusCode, e)
		return
	}

	user := domain.NewUserUpdate(uReq.Password, uReq.Name, uReq.Age)

	if err := uc.service.UpdateUserService(id, user); err != nil {
		logger.NewErrorLog(logger.JOURNEY["UPDATE_CONTROLLER"], logger.MESSAGE["ERROR"]["UPDATE"])
		c.JSON(err.StatusCode, err)
		return
	}

	var res adapters.TSuccessResponse
	res.Message = "User updated successfully."

	logger.NewInfoLog(logger.JOURNEY["UPDATE_CONTROLLER"], logger.MESSAGE["OK"]["UPDATED"])
	c.JSON(http.StatusOK, res)
}
