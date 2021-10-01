package models

import (
	"github.com/satori/go.uuid"
)

type Category struct {
	Base
	Name           string       `json:"name"`
	ShoppingListID uuid.UUID    `json:"shopping_list_id"`
	ShoppingList   ShoppingList `gorm:"foreignKey:ShoppingListID; references:ID"`
	Items          []Item
}
