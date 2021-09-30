package models

type ShoppingList struct {
  Base
  Name  string `json:"name" gorm:"type:varchar(100) not null;unique;index"`
}
