package goutil

func DefaultString(src, def string) string {
	if src == "" {
		return def
	}
	return src
}

func DefaultInt64(src, def int64) int64 {
	if src == 0 {
		return def
	}
	return src
}
