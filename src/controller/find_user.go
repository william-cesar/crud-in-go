package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/view/adapters"
)

func (uc *tUserController) FindUser(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		e := ierrors.NewBadRequestError()
		c.JSON(e.StatusCode, e)
		return
	}

	var u adapters.TUserEmailRequest

	if err = json.Unmarshal(body, &u); err != nil {
		e := ierrors.NewInternalError()
		c.JSON(e.StatusCode, e)
		return
	}

	user, uErr := uc.service.FindUserService(u.Email)

	if uErr != nil {
		c.JSON(uErr.StatusCode, uErr)
		return
	}

	c.JSON(http.StatusOK, adapters.ConvertDomainToResponse(user))
}
