package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/config/logger"
	"github.com/william-cesar/crud-in-go/src/view/adapters"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *tUserController) ActivateUser(c *gin.Context) {
	logger.NewInfoLog(logger.JOURNEY["ACTIVATE"], logger.MESSAGE["INIT"]["ACTIVATION"])

	id := c.Param("id")

	if _, err := primitive.ObjectIDFromHex(id); err != nil || strings.TrimSpace(id) == "" {
		logger.NewErrorLog(logger.JOURNEY["ACTIVATE"], logger.MESSAGE["ERROR"]["INVALID_ID"])
		e := ierrors.NewBadRequestError()
		c.JSON(e.StatusCode, e)
		return
	}

	if err := uc.service.ActivateUserService(id); err != nil {
		logger.NewErrorLog(logger.JOURNEY["ACTIVATE"], logger.MESSAGE["ERROR"]["ACTIVATION"])
		c.JSON(err.StatusCode, err)
		return
	}

	var res adapters.TSuccessResponse
	res.Message = "User activated successfully."

	logger.NewInfoLog(logger.JOURNEY["ACTIVATE"], logger.MESSAGE["OK"]["ACTIVE"])

	c.JSON(http.StatusOK, res)
}
