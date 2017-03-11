package goutil

import "testing"

// ShouldBeButIs
func ShouldBeButIs(t *testing.T, name string, expectedValue, value interface{}) {
	t.Errorf("%s should be %v but is %v", name, expectedValue, value)
}
