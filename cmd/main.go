package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"../app/routes"
	"../app/models"
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
	app.GetRoutes(s)
	s.Run(InitHost())


}
