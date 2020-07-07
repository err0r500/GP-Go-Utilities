package helpers

import (
	"github.com/VoodooTeam/GP-Go-Utilities/logger"
	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error, rawData ...interface{}) bool {
	if err != nil {
		logger.Error(err.Error())
		for _, raw := range rawData {
			logger.Info(raw)
		}

		switch e := err.(type) {
		case *HTTPError:
			c.JSON(e.Status, e)
		default:
			c.JSON(400, gin.H{"error": err.Error()})
		}
		return true
	}
	return false
}

func HandleErrorNotFound(c *gin.Context, err error) bool {
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return true
	}
	return false
}
