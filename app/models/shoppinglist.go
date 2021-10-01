package models

import (
	"encoding/json"

	util "../../utils/sql"
)

type ShoppingList struct {
	Base
	Name string `json:"name" gorm:"type:varchar(100) not null;unique;index"`
	Categories []Category
}


func GetAllShoppingListsAsTree()(output []byte, errMsg []byte){
	var shoppingList []ShoppingList
	if err:= DB.Preload("Categories").Find(&shoppingList).Error; err !=nil {
		errX := util.HandleSQLError(err)
		errMsg, _ = json.Marshal(errX)
		return nil, errMsg
	}
	output, _ = json.Marshal(shoppingList)
	return output, nil
}

func GetAllShoppingLists()(output []byte, errMsg []byte){
	var shoppingList []ShoppingList
	if err:= DB.Find(&shoppingList).Error; err !=nil {
		errX := util.HandleSQLError(err)
		errMsg, _ = json.Marshal(errX)
		return nil, errMsg
	}
	output, _ = json.Marshal(shoppingList)
	return output, nil
}
