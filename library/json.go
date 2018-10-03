package library

import (
	"encoding/json"
)

func ConvertToJSON(data interface{}) string{
	jsonResult, _ := json.Marshal(data)

	return string(jsonResult)
}