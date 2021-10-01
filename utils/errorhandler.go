package util

import (
	"encoding/json"
)

type SQLError struct {
Severity	string
Code	string
Message	string
Detail	string
Hint	string
Position	string
InternalPosition	string
InternalQuery	string
Where	string
SchemaName	string
TableName	string
ColumnName	string
DataTypeName	string
ConstraintName	string
File	string
Line	string
Routine	string
}

type Error struct {
	Detail string `json="detail"`
}

func HandleSQLError (err interface{}) (expError Error) {
		var sqlError SQLError
		byteError, _ := json.Marshal(err)
		json.Unmarshal(byteError, &sqlError)
		expError = Error{detail: sqlError.Detail}
		return expError
}

