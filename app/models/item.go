package models


type Category struct {
  ID     int   `json:"id" gorm:"primary_key"`
  Name  string `json:"name"`
}

type Item struct {
  ID     int   `json:"id" gorm:"primary_key"`
  Name  string `json:"name"`
  Bought bool `json:"bought"`
  CategoryID int `json:"category_id"`
  Category Category
}
