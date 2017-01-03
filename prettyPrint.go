package goutil

import "encoding/json"

// PrettyPrint a JSON
func PrettyPrint(x interface{}) string {
	b, _ := json.MarshalIndent(x, "", "  ")
	return string(b)
}
