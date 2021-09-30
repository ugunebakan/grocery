package models

import (
	"github.com/satori/go.uuid"
)

type Item struct {
  Base
  Name  string `json:"name" gorm:"type:varchar(100) not null;unique;index"`
  Bought bool `json:"bought"`
  CategoryID uuid.UUID `json:"category_id"`
  Category Category `gorm:"foreignKey:CategoryID; references:ID"`
}
