package controller

import (
	"net/http"

	"../models"
	"../../utils"
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
	if err := models.DB.Create(&itemPost).Error; err != nil {
		umut := util.HandleSQLError(err)
		c.JSON(http.StatusBadRequest, umut)
	} else {
		c.JSON(http.StatusCreated, &itemPost)
	}
}
