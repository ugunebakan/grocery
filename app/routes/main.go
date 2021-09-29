package app

import (
	"github.com/gin-gonic/gin"
	//"fmt"
)

//var (
	//app = gin.Default()
//)

//func Run(hostname string) {
	//getRoutes()
	//app.Run(hostname)
//}

func GetRoutes(app *gin.Engine) {
	v1 := app.Group("/v1")
	addPingRoutes(v1)
}
