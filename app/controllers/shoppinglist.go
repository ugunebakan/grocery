package controller

import (
	"log"
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

type ShoppingList struct {
	ID         string
	Name       string
	Categories *[]Category
}

func ShoppingListGET(c *gin.Context) {
	var shoppingList []ShoppingList
	models.DB.Preload("Categories.Items").Preload("Categories").Find(&shoppingList)
	c.JSON(http.StatusOK, shoppingList)
}

type ShoppingListPaylod struct {
	Name string
}

func ShoppingListPOST(c *gin.Context) {
	var shoppingList ShoppingListPaylod
	err := c.BindJSON(&shoppingList)
	if err != nil {
		log.Fatal(err)
	}
	shoppingListPost := models.ShoppingList{Name: shoppingList.Name}
	models.DB.Create(&shoppingListPost)
	c.JSON(http.StatusCreated, &shoppingListPost)
}
