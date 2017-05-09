package goutil

import (
	"testing"
)

func TestCypherSet(t *testing.T) {

	props := map[string]interface{}{}
	cs := NewCypherSets("person", props)

	name := "Jean"
	var agePtr *float64
	agePtr = nil

	cs.
		AddString("personName", "name", &name).
		AddFloat64("personAge", "age", agePtr)

	if _, exists := cs.Props["personAge"]; exists {
		t.Errorf("Should not have %s in props \n", "personAge")
	}

	if _, exists := cs.Props["name"]; !exists {
		t.Errorf("Should have %s in props \n", "name")
	}

}
