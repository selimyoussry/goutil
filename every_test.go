package goutil

import "testing"

func TestAny(t *testing.T) {

	if !Any("a", "b", "") {
		t.Errorf("Any, correct response is true")
	}

}

func TestEvery(t *testing.T) {

	if Every("a", "", "", "", "d") {
		t.Errorf("Every should be false")
	}

}

func TestIsIn(t *testing.T) {

	if IsIn("hello", "helloz", "", "hhee") {
		t.Errorf("IsIn: hello is not in the list")
	}

	if !IsIn("hello", "helloz", "", "hello", "hhee") {
		t.Errorf("IsIn: hello is in the list")
	}

}
