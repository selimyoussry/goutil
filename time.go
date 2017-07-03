package goutil

import (
	"time"
)

func UnixTimestampMilliseconds(t time.Time) int64 {

	nanoToMilli := int64(1000 * 1000)
	timestampNS := t.UnixNano()
	nanoSeconds := timestampNS % nanoToMilli

	timestampMS := (timestampNS - nanoSeconds) / nanoToMilli

	return timestampMS
}
