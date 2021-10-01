package controller

import (
	responseUtil "../../utils/response"
	"../models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ShoppingListGET(c *gin.Context) {
	data, err := models.GetAllShoppingLists()
	responseUtil.ResponseHandler(c, data, err, 200)
}

type ShoppingListPaylod struct {
	Name string `json:"name" binding:"required"`
}

func ShoppingListPOST(c *gin.Context) {
	var sl ShoppingListPaylod
	c.BindJSON(&sl)

	m := models.ShoppingList{
		Name: sl.Name,
	}

	data, err := m.CreateShoppingList()
	responseUtil.ResponseHandler(c, data, err, 201)
}

func ShoppingListPUT(c *gin.Context) {
	var sl ShoppingListPaylod
	id := c.Param("id")
	fmt.Printf("%T", id)
	fmt.Print("\n\n")
	c.BindJSON(&sl)

	m := models.ShoppingList{
		Name: sl.Name,
	}

	data, err := m.UpdateShoppingList(id)
	responseUtil.ResponseHandler(c, data, err, 201)
}
