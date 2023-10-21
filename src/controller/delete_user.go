package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/view/adapters"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *tUserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if _, err := primitive.ObjectIDFromHex(id); err != nil || strings.TrimSpace(id) == "" {
		e := ierrors.NewBadRequestError()
		c.JSON(e.StatusCode, e)
		return
	}

	if err := uc.service.DeleteUserService(id); err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	var res adapters.TSuccessResponse
	res.Message = "User deleted successfully."

	c.JSON(http.StatusOK, res)
}
