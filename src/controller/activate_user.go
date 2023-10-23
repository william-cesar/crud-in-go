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

// Activate				godoc
// @Summary 			Activate users
// @Description 	Activate user
// @Produce 			application/json
// @Tags 					journey
// @Param 				id path string true "id"
// @Success 			200 {object} adapters.TSuccessResponse
// @Failure				400 {object} ierrors.TError
// @Router 				/activate/{id}  [get]
func (uc *tUserController) ActivateUser(c *gin.Context) {
	logger.NewInfoLog(logger.JOURNEY["ACTIVATE_CONTROLLER"], logger.MESSAGE["INIT"]["ACTIVATION"])

	id := c.Param("id")

	if _, err := primitive.ObjectIDFromHex(id); err != nil || strings.TrimSpace(id) == "" {
		logger.NewErrorLog(logger.JOURNEY["ACTIVATE_CONTROLLER"], logger.MESSAGE["ERROR"]["INVALID_ID"], err)
		e := ierrors.NewBadRequestError()
		c.JSON(e.StatusCode, e)
		return
	}

	if err := uc.service.ActivateUserService(id); err != nil {
		logger.NewErrorLog(logger.JOURNEY["ACTIVATE_CONTROLLER"], logger.MESSAGE["ERROR"]["ACTIVATION"], err)
		c.JSON(err.StatusCode, err)
		return
	}

	var res adapters.TSuccessResponse
	res.Message = "User activated successfully."

	logger.NewInfoLog(logger.JOURNEY["ACTIVATE_CONTROLLER"], logger.MESSAGE["OK"]["ACTIVE"])
	c.JSON(http.StatusOK, res)
}
