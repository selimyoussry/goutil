package goutil

import "github.com/hippoai/goerr"

func ErrDefaultMustBeEven() error {
	return goerr.NewS(ERR_DEFAULT_MUST_BE_EVEN)
}
