package goutil

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {

	layout := "2006-01-02T15:04:05 -07:00"
	datetimeStr := "2010-02-16T11:12:13 +00:00"

	tObject, _ := time.Parse(layout, datetimeStr)
	timestampMS := UnixTimestampMilliseconds(tObject)

	expected := int64(1266318733000)

	if timestampMS != expected {
		t.Errorf("Wrong timestamp, got %v, expected %v", timestampMS, expected)
	}
}
