package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/config/logger"
	"github.com/william-cesar/crud-in-go/src/controller/handlers"
)

func Auth(c *gin.Context) {
	logger.NewInfoLog(logger.JOURNEY["AUTH"], logger.MESSAGE["INIT"]["AUTH"])

	token := c.GetHeader("Authorization")

	if token == "" {
		logger.NewInfoLog(logger.JOURNEY["AUTH"], logger.MESSAGE["ERROR"]["NO_TOKEN"])
		err := ierrors.NewUnauthorizedError()
		c.AbortWithStatusJSON(err.StatusCode, err)
	}

	err := handlers.VerifyToken(token)

	if err != nil {
		logger.NewInfoLog(logger.JOURNEY["AUTH"], logger.MESSAGE["ERROR"]["INVALID_TOKEN"])
		c.AbortWithStatusJSON(err.StatusCode, err)
	}

	logger.NewInfoLog(logger.JOURNEY["AUTH"], logger.MESSAGE["OK"]["AUTH"])
	c.Next()
}
