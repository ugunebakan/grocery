package router

import (
	"github.com/gin-gonic/gin"
	"../controllers"
)

func addShoppingListRoutes(rg *gin.RouterGroup) {
	item := rg.Group("/shoppinglist")
	item.GET("/", controller.ShoppingListGET)
	//item.POST("/", controller.ItemPOST)
	//item.GET("/:id", controller.CategoryRETRIEVE)
}
