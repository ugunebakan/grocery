package util

import (
	"github.com/gin-gonic/gin"
)

func ResponseHandler(c *gin.Context, data []byte, err []byte, responseCode int) {
	if err != nil {
		if responseCode == 0 {
			responseCode = 400
		}
		c.Data(responseCode, "application/json", err)
	} else {
		c.Data(responseCode, "application/json", data)
	}
}
