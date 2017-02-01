package goutil

// GetKeysFromMSI returns the keys of a map[string]interface
func GetKeysFromMSI(m map[string]interface{}) []string {
	keys := make([]string, len(m))
	i := 0
	for key, _ := range m {
		keys[i] = key
		i += 1
	}
	return keys
}

func SameKeys(a, b map[string]interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	var exists bool
	for keyA, _ := range a {
		_, exists = b[keyA]
		if !exists {
			return false
		}
	}

	return true
}
