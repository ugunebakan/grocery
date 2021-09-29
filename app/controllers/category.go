package controller

import (
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

func CategoryController(c *gin.Context){
	categoryPost := models.Category{Name: "input.Title"}
	models.DB.Create(&categoryPost)


	var categoryGet []models.Category
	models.DB.Find(&categoryGet)
	c.JSON(http.StatusOK, gin.H{"data": categoryGet})
}
