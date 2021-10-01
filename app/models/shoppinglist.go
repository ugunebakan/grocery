package models

import (
	"encoding/json"
	//"errors"
	"html"
	//"fmt"
	"github.com/satori/go.uuid"

	util "../../utils/sql"
	"github.com/microcosm-cc/bluemonday"
)

type ShoppingList struct {
	Base
	Name       string `json:"name" gorm:"type:varchar(100) not null;unique;index"`
	Categories []Category
}

func GetAllShoppingListsAsTree() (output []byte, errMsg []byte) {
	var shoppingList []ShoppingList
	if err := DB.Preload("Categories").Find(&shoppingList).Error; err != nil {
		errX := util.HandleSQLError(err)
		errMsg, _ = json.Marshal(errX)
		return nil, errMsg
	}
	output, _ = json.Marshal(shoppingList)
	return output, nil
}

func GetAllShoppingLists() (output []byte, errMsg []byte) {
	var shoppingList []ShoppingList
	if err := DB.Find(&shoppingList).Error; err != nil {
		errX := util.HandleSQLError(err)
		errMsg, _ = json.Marshal(errX)
		return nil, errMsg
	}
	output, _ = json.Marshal(shoppingList)
	return output, nil
}

type ShoppingListError struct {
	Name []string `json:"name"`
}

func (sl *ShoppingList) Validate() (errMsg []byte) {
	var errorModel ShoppingListError
	var errorList []string

	p := bluemonday.StrictPolicy()
	if sl.Name == "" {
		errorList = append(errorList, "Name cannot be empty")
	}

	if len(errorList) == 0 {
		return nil
	}
	errorModel = ShoppingListError{Name: errorList}
	errMsg, _ = json.Marshal(&errorModel)

	sl.Name = html.UnescapeString(p.Sanitize(sl.Name))
	return errMsg
}

func (sl *ShoppingList) CreateShoppingList() (output []byte, errMsg []byte) {
	if err := sl.Validate(); err != nil {
		return nil, err
	}
	if err := DB.Create(&sl).Error; err != nil {
		errX := util.HandleSQLError(err)
		errMsg, _ = json.Marshal(errX)
		return nil, errMsg
	}
	output, _ = json.Marshal(&sl)
	return output, nil
}

func (sl *ShoppingList) UpdateShoppingList(id string) (output []byte, errMsg []byte) {
	if err := sl.Validate(); err != nil {
		return nil, err
	}

	sl.ID, _ = uuid.FromString(id)
	if err := DB.Where("id = ?", id).First(&sl).Error; err != nil {
		errX := util.HandleSQLError(err)
		errMsg, _ = json.Marshal(errX)
		return nil, errMsg
	}

	if err := DB.Updates(&sl).Error; err != nil {
		errX := util.HandleSQLError(err)
		errMsg, _ = json.Marshal(errX)
		return nil, errMsg
	}

	output, _ = json.Marshal(&sl)
	return output, nil
}
