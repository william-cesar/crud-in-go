package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/config/validations"
	"github.com/william-cesar/crud-in-go/src/model/domain"
	"github.com/william-cesar/crud-in-go/src/view/adapters"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *tUserController) UpdateUser(c *gin.Context) {
	var uReq adapters.TUserUpdateRequest

	if err := c.ShouldBindJSON(&uReq); err != nil {
		reqErr := validations.ValidateUserError(err)
		c.JSON(reqErr.StatusCode, reqErr)
		return
	}

	id := c.Param("id")

	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		e := ierrors.NewBadRequestError()
		c.JSON(e.StatusCode, e)
		return
	}

	user := domain.NewUserUpdate(uReq.Password, uReq.Name, uReq.Age)

	updatedUser, err := uc.service.UpdateUserService(id, user)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, adapters.ConvertDomainToResponse(updatedUser))
}
