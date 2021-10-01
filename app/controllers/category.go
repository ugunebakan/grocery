package controller

import (
	"log"
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

// Item...
type Item struct {
	ID         string
	Name       string
	Bought     bool
	CategoryID string
}

//Category..
type Category struct {
	ID             string
	Name           string
	ShoppingListID string
	Items          []Item
}

func CategoryGET(c *gin.Context) {
	var categoryGet []Category
	models.DB.Preload("Items").Find(&categoryGet)
	c.JSON(http.StatusOK, categoryGet)
}

func CategoryPOST(c *gin.Context) {
	var categoryInput models.Category
	err := c.BindJSON(&categoryInput)
	if err != nil {
		log.Fatal(err)
	}
	categoryPost := models.Category{Name: categoryInput.Name, ShoppingListID: categoryInput.ShoppingListID}
	models.DB.Create(&categoryPost)
	c.JSON(http.StatusCreated, categoryPost)
}

func CategoryRETRIEVE(c *gin.Context) {
	var categoryRetrieve models.Category
	models.DB.Where("id = ?", c.Param("id")).First(&categoryRetrieve)
	c.JSON(http.StatusOK, categoryRetrieve)
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Category
// @Router /example/helloworld [get]
