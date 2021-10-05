package router

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

func addShoppingListRoutes(rg *gin.RouterGroup) {
	shoppinglist := rg.Group("/shoppinglist")
	shoppinglist.GET("/", controller.ShoppingListGET)
	shoppinglist.GET("/:id", controller.ShoppingListOneGET)
	shoppinglist.POST("/", controller.ShoppingListPOST)
	shoppinglist.PUT("/:id", controller.ShoppingListPUT)
}
