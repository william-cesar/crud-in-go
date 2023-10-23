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

// DeleteUser			godoc
// @Summary 			Delete users
// @Description 	Delete user
// @Produce 			application/json
// @Param 				id path string true "id"
// @Tags 					user
// @Success 			200 {object} adapters.TSuccessResponse
// @Failure				400 {object} ierrors.TError
// @Router 				/user/{id}  [delete]
// @Param 				Authorization header string true "Access token" default(Bearer <Add access token here>)
func (uc *tUserController) DeleteUser(c *gin.Context) {
	logger.NewInfoLog(logger.JOURNEY["DELETE_CONTROLLER"], logger.MESSAGE["INIT"]["DELETION"])

	id := c.Param("id")

	if _, err := primitive.ObjectIDFromHex(id); err != nil || strings.TrimSpace(id) == "" {
		logger.NewErrorLog(logger.JOURNEY["DELETE_CONTROLLER"], logger.MESSAGE["ERROR"]["INVALID_ID"], err)
		e := ierrors.NewBadRequestError()
		c.JSON(e.StatusCode, e)
		return
	}

	if err := uc.service.DeleteUserService(id); err != nil {
		logger.NewErrorLog(logger.JOURNEY["DELETE_CONTROLLER"], logger.MESSAGE["ERROR"]["DELETION"], err)
		c.JSON(err.StatusCode, err)
		return
	}

	var res adapters.TSuccessResponse
	res.Message = "User deleted successfully."

	logger.NewInfoLog(logger.JOURNEY["DELETE_CONTROLLER"], logger.MESSAGE["OK"]["DELETED"])
	c.JSON(http.StatusOK, res)
}
