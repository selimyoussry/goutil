package goutil

import (
	"testing"
)

func TestNilCheck(t *testing.T) {

	var s string
	s = "hello World"

	var b bool
	b = true

	var f float64
	f = 1.3

	var arrS []string
	arrS = []string{"a", "b"}

	var nilS *string
	var nilF *float32

	var isNil bool
	var err error

	isNil, err = NilCheck(&s)
	if isNil {
		t.Errorf("Should be false - 1")
	}

	isNil, err = NilCheck(nilS)
	if !isNil {
		t.Errorf("Should be true - 2")
	}

	isNil, err = NilCheck(&b)
	if isNil {
		t.Errorf("Should be false - 3")
	}

	isNil, err = NilCheck(&arrS)
	if err == nil {
		t.Errorf("Should not know - 4")
	}

	isNil, err = NilCheck(&f)
	if isNil {
		t.Errorf("Should be false - 5")
	}

	isNil, err = NilCheck(nilF)
	if !isNil {
		t.Errorf("Should be true - 6")
	}

	isNil, err = NilCheck(nil)
	if !isNil {
		t.Errorf("Should be true - 7")
	}

}
