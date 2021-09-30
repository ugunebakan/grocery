package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
        ginSwagger "github.com/swaggo/gin-swagger"

	"../app/routes"
	"../app/models"
	_ "../docs"	
)

func InitHost()(string){
	port := os.Getenv("PORT")
	if port == "" {
		 port = "3000"
	}
	hostname := os.Getenv("HOSTNAME")
	if hostname == "" {
		hostname = "localhost"
	}

	return hostname + ":" + port
}


// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := os.Getenv("ENVIRONMENT")
	if "" == env {
		env = "development"
	}

	models.ConnectDataBase()
	s := gin.Default()
	router.GetRoutes(s)
	url := ginSwagger.URL("http://localhost:8000/swagger/doc.json") 
	s.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))
	s.Run(InitHost())


}
