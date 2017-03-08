package goutil

import (
	"fmt"
	"regexp"
)

// ExtractPattern will returns a slice of parsed parenthisized values
// found in the string s, given a regex pattern
// Each of these values is stored in a map
func RegexExtractPattern(s, pattern string) [](map[string]string) {
	re := regexp.MustCompile(pattern)
	results := re.FindAllString(s, -1)

	m := [](map[string]string){}
	for _, result := range results {
		parsed := map[string]string{}

		for _, key := range re.SubexpNames() {
			if key == "" {
				continue
			}
			parsed[key] = re.ReplaceAllString(result, fmt.Sprintf("${%s}", key))
		}
		m = append(m, parsed)
	}
	return m
}
