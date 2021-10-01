package controller

import (
	"log"
	"net/http"

	"../../utils/sql"
	responseUtil "../../utils/response"
	"../models"
	"github.com/gin-gonic/gin"
)

func ItemGET(c *gin.Context) {
	items, err := models.GetAllItems()
	responseUtil.ResponseHandler(c, items, err, 200)
}

func ItemPOST(c *gin.Context) {
	var itemInput models.Item
	err := c.BindJSON(&itemInput)
	if err != nil {
		log.Fatal(err)
	}
	itemPost := models.Item{Name: itemInput.Name, CategoryID: itemInput.CategoryID}
	if err := models.DB.Create(&itemPost).Error; err != nil {
		umut := util.HandleSQLError(err)
		c.JSON(http.StatusBadRequest, umut)
	} else {
		c.JSON(http.StatusCreated, &itemPost)
	}
}
