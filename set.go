package goutil

type SetString map[string]bool

func NewSetString() *SetString {
	return &SetString{}
}

func (s *SetString) Add(elts ...string) *SetString {
	for _, elt := range elts {
		(*s)[elt] = true
	}
	return s
}

func (s *SetString) Remove(elt string) *SetString {
	delete(*s, elt)
	return s
}

func (s *SetString) Exists(elt string) bool {
	_, exists := (*s)[elt]
	return exists
}

func SetString_ToSlice(s *SetString) []string {
	ss := []string{}
	for elt, _ := range *s {
		ss = append(ss, elt)
	}
	return ss
}

func Slice_ToSetString(elts []string) *SetString {
	s := NewSetString()
	s.Add(elts...)
	return s
}
