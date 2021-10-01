package util

import (
	"github.com/gin-gonic/gin"
)

func ResponseHandler(c *gin.Context, data []byte, err []byte, responseCode int) {
	if err != nil {
		c.Data(400, "application/json", err)
	} else {
		c.Data(responseCode, "application/json", data)
	}
}
