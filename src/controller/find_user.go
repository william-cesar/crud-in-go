package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/config/logger"
	"github.com/william-cesar/crud-in-go/src/view/adapters"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *tUserController) FindUserByEmail(c *gin.Context) {
	logger.NewInfoLog(logger.JOURNEY["FIND_CONTROLLER"], logger.MESSAGE["INIT"]["FIND_EMAIL"])

	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		logger.NewErrorLog(logger.JOURNEY["FIND_CONTROLLER"], logger.MESSAGE["ERROR"]["REQ_BODY"])
		e := ierrors.NewBadRequestError()
		c.JSON(e.StatusCode, e)
		return
	}

	var u adapters.TUserEmailRequest

	if err = json.Unmarshal(body, &u); err != nil {
		logger.NewErrorLog(logger.JOURNEY["FIND_CONTROLLER"], logger.MESSAGE["ERROR"]["UNMARSHAL"])
		e := ierrors.NewInternalError()
		c.JSON(e.StatusCode, e)
		return
	}

	user, uErr := uc.service.FindUserByEmailService(u.Email)

	if uErr != nil {
		logger.NewErrorLog(logger.JOURNEY["FIND_CONTROLLER"], logger.MESSAGE["ERROR"]["NO_USER"])
		c.JSON(uErr.StatusCode, uErr)
		return
	}

	logger.NewInfoLog(logger.JOURNEY["FIND_CONTROLLER"], logger.MESSAGE["OK"]["FOUND"])
	c.JSON(http.StatusOK, adapters.ConvertDomainToResponse(user))
}

func (uc *tUserController) FindUserById(c *gin.Context) {
	logger.NewInfoLog(logger.JOURNEY["FIND_CONTROLLER"], logger.MESSAGE["INIT"]["FIND_ID"])

	id := c.Param("id")

	if _, err := primitive.ObjectIDFromHex(id); err != nil || strings.TrimSpace(id) == "" {
		logger.NewErrorLog(logger.JOURNEY["FIND_CONTROLLER"], logger.MESSAGE["ERROR"]["INVALID_ID"])
		e := ierrors.NewBadRequestError()
		c.JSON(e.StatusCode, e)
		return
	}

	user, uErr := uc.service.FindUserByIdService(id)

	if uErr != nil {
		logger.NewErrorLog(logger.JOURNEY["FIND_CONTROLLER"], logger.MESSAGE["ERROR"]["NO_USER"])
		c.JSON(uErr.StatusCode, uErr)
		return
	}

	logger.NewInfoLog(logger.JOURNEY["FIND_CONTROLLER"], logger.MESSAGE["OK"]["FOUND"])
	c.JSON(http.StatusOK, adapters.ConvertDomainToResponse(user))
}
