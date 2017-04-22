package goutil

import "strings"

// URLJoin takes a bunch of urls and concatenate them
func URLJoin(ss ...string) string {
	parts := []string{}

	for _, s := range ss {

		// Skip empty strings
		if s == "" {
			continue
		}

		// Concatenate properly
		parts = append(parts, urlJoinRemoveExtremeSlashes(s))
	}

	return strings.Join(parts, "/")
}

func urlJoinRemoveExtremeSlashes(s string) string {
	if s == "" {
		return s
	}

	firstLetter := s[:1]
	lastLetter := s[(len(s) - 1):]

	newS := s
	if firstLetter == "/" {
		newS = newS[1:]
	}
	if lastLetter == "/" {
		newS = newS[:(len(newS) - 1)]
	}

	return newS
}
