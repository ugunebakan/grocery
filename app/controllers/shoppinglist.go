package controller

import (
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

type ShoppingList struct {
	ID string
	Name string
	Categories []Category
}


func ShoppingListGET(c *gin.Context){
	var shoppingList ShoppingList
	models.DB.Preload("Categories.Items").Preload("Categories").Find(&shoppingList)
	c.JSON(http.StatusOK, shoppingList)
}

//func ItemPOST(c *gin.Context){
	//var itemInput models.Item
	//c.BindJSON(&itemInput)
	//itemPost := models.Item{Name: itemInput.Name, CategoryID: itemInput.CategoryID}
	//models.DB.Create(&itemPost)
	//c.JSON(http.StatusCreated, &itemPost)
//}
