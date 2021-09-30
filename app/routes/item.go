package router

import (
	"github.com/gin-gonic/gin"
	"../controllers"
)

func addItemRoutes(rg *gin.RouterGroup) {
	item := rg.Group("/item")
	item.GET("/", controller.ItemGET)
	item.POST("/", controller.ItemPOST)
	//item.GET("/:id", controller.CategoryRETRIEVE)
}
