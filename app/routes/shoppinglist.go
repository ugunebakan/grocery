package router

import (
	"github.com/gin-gonic/gin"
	"../controllers"
)

func addShoppingListRoutes(rg *gin.RouterGroup) {
	shoppinglist := rg.Group("/shoppinglist")
	shoppinglist.GET("/", controller.ShoppingListGET)
	shoppinglist.POST("/", controller.ShoppingListPOST)
	//item.GET("/:id", controller.CategoryRETRIEVE)
}
