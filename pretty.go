package goutil

import "encoding/json"

// Pretty makes a thing pretty for print a JSON
func Pretty(x interface{}) string {
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		return "__err__"
	}
	return string(b)
}
