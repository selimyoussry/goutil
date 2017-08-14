package goutil

import (
	"fmt"
	"os"
)

func Log(s string, i ...interface{}) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf("%s \n", s), i...)
}
