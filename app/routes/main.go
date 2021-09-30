package router

import (
	"github.com/gin-gonic/gin"
)


func GetRoutes(app *gin.Engine) {
	api := app.Group("/api")
	addApiRoutes(api)
}

func addApiRoutes(rt *gin.RouterGroup) {
	v1 := rt.Group("/v1")
	addCategoryRoutes(v1)
	addItemRoutes(v1)
}
