package controller

import (
	"log"
	"net/http"

	"../models"
	responseUtil "../../utils/response"
	"github.com/gin-gonic/gin"
)


func ShoppingListGET(c *gin.Context) {
	data, err := models.GetAllShoppingLists()
	responseUtil.ResponseHandler(c, data, err, 200)
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
