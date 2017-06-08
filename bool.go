package goutil

// SumBool returns the number of true booleans
func SumBool(bs ...bool) int {
	sum := 0
	for _, b := range bs {
		if b {
			sum++
		}
	}
	return sum
}
