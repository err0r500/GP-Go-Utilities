package helpers

import (
	bugsnag "github.com/bugsnag/bugsnag-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func HandleError(c *gin.Context, err error) bool {
	if err != nil {
		log.Error(err.Error())
		bugsnag.Notify(err, c.Request.Context())
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
