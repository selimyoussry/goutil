package goutil

// Every returns true if any are equal to ""
func Any(ss ...string) bool {
	for _, s := range ss {
		if s == "" {
			return true
		}
	}
	return false
}

// Every returns true if all are equal to ""
func Every(ss ...string) bool {
	for _, s := range ss {
		if s != "" {
			return false
		}
	}
	return true
}

// IsIn returns true iff s is in ss
func IsIn(a string, ss ...string) bool {
	for _, s := range ss {
		if s == a {
			return true
		}
	}
	return false
}
