package controller

import (
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)


func ItemGET(c *gin.Context){
	var ItemGet []models.Item
	models.DB.Find(&ItemGet)
	c.JSON(http.StatusOK, ItemGet)
}

func ItemPOST(c *gin.Context){
	var itemInput models.Item
	c.BindJSON(&itemInput)
	itemPost := models.Item{Name: itemInput.Name, CategoryID: itemInput.CategoryID}
	models.DB.Create(&itemPost)
	c.JSON(http.StatusCreated, &itemPost)
}
