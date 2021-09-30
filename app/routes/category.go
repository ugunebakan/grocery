package router

import (
	"github.com/gin-gonic/gin"
	"../controllers"
)

func addCategoryRoutes(rg *gin.RouterGroup) {
	category := rg.Group("/category")
	category.GET("/", controller.CategoryGET)
	category.POST("/", controller.CategoryPOST)
	category.GET("/:id", controller.CategoryRETRIEVE)
}
