package helper

import (
	"encoding/json"
	"net/http"
)

// ParseJSON ...
func ParseJSON(res *http.Response) map[string]interface{} {
	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)
	return result
}
