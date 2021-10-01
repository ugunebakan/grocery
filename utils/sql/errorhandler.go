package util

import (
	"encoding/json"
	"log"
)

type SQLError struct {
	Severity         string
	Code             string
	Message          string
	Detail           string
	Hint             string
	Position         string
	InternalPosition string
	InternalQuery    string
	Where            string
	SchemaName       string
	TableName        string
	ColumnName       string
	DataTypeName     string
	ConstraintName   string
	File             string
	Line             string
	Routine          string
}

type Error struct {
	Detail string `json:"detail"`
}

func HandleSQLError(errData interface{}) (expError Error) {
	var sqlError SQLError
	var byteError []byte
	byteError, _ = json.Marshal(errData)
	err := json.Unmarshal(byteError, &sqlError)
	if err != nil {
		log.Print(err)
	}
	expError = Error{Detail: sqlError.Detail}
	return expError
}
