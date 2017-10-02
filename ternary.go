package goutil

import "time"

// Implements cond ? _true : _false
func TerString(cond bool, _true, _false string) string {
	if cond {
		return _true
	}

	return _false
}

// Implements cond ? _true : _false
func TerBytes(cond bool, _true, _false []byte) []byte {
	if cond {
		return _true
	}

	return _false
}

// Implements cond ? _true : _false
func TerTime(cond bool, _true, _false time.Time) time.Time {
	if cond {
		return _true
	}

	return _false
}

// Implements cond ? _true : _false
func TerInt(cond bool, _true, _false int) int {
	if cond {
		return _true
	}

	return _false
}

// Implements cond ? _true : _false
func TerBool(cond bool, _true, _false bool) bool {
	if cond {
		return _true
	}

	return _false
}

// Implements cond ? _true : _false
func TerFloat64(cond bool, _true, _false float64) float64 {
	if cond {
		return _true
	}

	return _false
}

// Implements cond ? _true : _false
func TerFloat32(cond bool, _true, _false float32) float32 {
	if cond {
		return _true
	}

	return _false
}
