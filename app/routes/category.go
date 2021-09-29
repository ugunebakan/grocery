package app

import (
	"github.com/gin-gonic/gin"
	"../controllers"
)

func addPingRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/ping")
	ping.GET("/", controller.CategoryController)
}
