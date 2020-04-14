package utilities

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func HandleError(c *gin.Context, err error) bool {
	if err != nil {
		log.Error(err.Error())
		switch e := err.(type) {
		case *HTTPError:
			c.JSON(e.Status, e)
		default:
			c.JSON(400, gin.H{"error": err.Error()})
			return true
		}
	}
	return false
}
