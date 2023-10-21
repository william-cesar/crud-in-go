package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/controller/handlers"
)

func Authorized(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == "" {
		err := ierrors.NewUnauthorizedError()
		c.AbortWithStatusJSON(err.StatusCode, err)
	}

	err := handlers.VerifyToken(token)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, err)
	}

	c.Next()
}
