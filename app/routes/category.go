package router

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

func addCategoryRoutes(rg *gin.RouterGroup) {
	category := rg.Group("/category")
	category.GET("/", controller.CategoryGET)
	category.POST("/", controller.CategoryPOST)
	category.GET("/:id", controller.CategoryRETRIEVE)
}
