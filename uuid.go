package goutil

import "github.com/twinj/uuid"

func UuidV4() string {
	return uuid.NewV4().String()
}
