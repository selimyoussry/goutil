package goutil

import (
	"time"
)

func MillisecondsToTime(timestamp_ms int64) time.Time {

	milliToNano := int64(1000 * 1000)
	milliseconds := timestamp_ms % 1000
	nanoseconds := milliseconds * milliToNano
	seconds := (timestamp_ms - milliseconds) / 1000

	return time.Unix(seconds, nanoseconds).UTC()
}

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
