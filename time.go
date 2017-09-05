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

func TimeCopy(t time.Time) time.Time {
	formattedTime := t.Format(time.RFC3339Nano)
	newT, _ := time.Parse(time.RFC3339Nano, formattedTime)
	return newT
}
