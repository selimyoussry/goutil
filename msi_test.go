package goutil

import "testing"

func TestGetKeysFromMSI(t *testing.T) {

	m := map[string]interface{}{
		"a": 1,
		"b": "hello",
		"z": true,
	}

	keys := []string{"a", "b", "z"}

	fKeys := GetKeysFromMSI(m)

	for _, key := range fKeys {
		if !IsIn(key, keys...) {
			t.Errorf("%s is not is keys, but should be", key)
		}
	}

}
