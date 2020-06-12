package helpers

import (
	bugsnag "github.com/bugsnag/bugsnag-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func HandleError(c *gin.Context, err error, rawData ...interface{}) bool {
	if err != nil {
		log.Error(err.Error())

		metadata := bugsnag.MetaData{}
		for _, raw := range rawData {
			metadata.AddStruct("metadata", raw)
		}
		user := bugsnag.User{}
		if userID := c.GetString("uid"); userID != "" {
			user.Id = userID
		}
		bugsnag.Notify(err, c.Request.Context(), metadata, user)

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
