package utils

import (
	"encoding/json"
	"fmt"
)

func JsonMarshal(input interface{}) string {
	jsonString, _ := json.Marshal(input)
	return fmt.Sprintf("%s\n", string(jsonString))
}
