package goutil

import (
	"crypto/md5"
	"fmt"
)

// MD5 helper function
func MD5(input string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(input)))
}
