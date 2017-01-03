package goutil

import "strings"

// Keify formats a string for a key
func Keify(s string) string {

	var tmp string
	tmp = strings.ToLower(s)
	tmp = strings.Replace(tmp, " ", "_", -1)
	tmp = strings.Replace(tmp, "'", "_", -1)

	return tmp

}
