package models

import (
	"../../utils/sql"
	"encoding/json"
	"github.com/satori/go.uuid"
)

type Item struct {
	Base
	Name       string    `json:"name" gorm:"type:varchar(100) not null;unique;index"`
	Bought     bool      `json:"bought"`
	CategoryID uuid.UUID `json:"category_id"`
	Category   Category  `gorm:"foreignKey:CategoryID; references:ID"`
}

func GetAllItems() (output []byte, errMsg []byte) {
	var items []Item
	if err := DB.Find(&items).Error; err != nil {
		errX := util.HandleSQLError(err)
		errMsg, _ = json.Marshal(errX)
		return nil, errMsg
	}
	output, _ = json.Marshal(items)
	return output, nil
}
