package goutil

import (
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Keify formats a string for a key
func Keify(s string) string {

	var tmp string
	tmp = strings.ToLower(s)
	tmp = strings.TrimSpace(tmp)
	tmp = strings.Replace(tmp, "  ", " ", -1)
	tmp = strings.Replace(tmp, " ", "_", -1)
	tmp = strings.Replace(tmp, "'", "_", -1)
	tmp = strings.Replace(tmp, "(", "_", -1)
	tmp = strings.Replace(tmp, ")", "_", -1)

	return tmp

}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func KeifyNFC(s string) (string, error) {
	keifiedS := Keify(s)

	// Input bytes
	b := make([]byte, len(keifiedS))

	// Transform here
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	_, _, err := t.Transform(b, []byte(keifiedS), true)
	if err != nil {
		return keifiedS, err
	}

	return string(b), nil
}
