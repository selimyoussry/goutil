package goutil

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
)

// URL wraps a url string with its GET parameters
type URL struct {
	URL           string
	GetParameters map[string]interface{}
}

// NewURL instanciates
func NewURL(url string) *URL {
	return &URL{
		URL:           url,
		GetParameters: map[string]interface{}{},
	}
}

// URLJoin takes a bunch of urls and concatenate them
func URLJoin(ss ...string) *URL {
	parts := []string{}

	for _, s := range ss {

		// Skip empty strings
		if s == "" {
			continue
		}

		// Concatenate properly
		parts = append(parts, urlJoinRemoveExtremeSlashes(s))
	}

	return NewURL(strings.Join(parts, "/"))
}

// SetURLParameter sets one parameters
func (u *URL) SetURLJSONParameter(key string, value interface{}) *URL {
	b, err := json.Marshal(value)
	if err != nil {
		log.Fatalf("%v", Pretty(err))
	}

	u.GetParameters[key] = string(b)
	return u
}

// SetURLParameter sets one parameters
func (u *URL) SetURLParameter(key string, value interface{}) *URL {
	u.GetParameters[key] = value
	return u
}

// SetURLParameters sets one parameters
func (u *URL) SetURLParameters(params map[string]interface{}) *URL {
	for key, value := range params {
		u.GetParameters[key] = value
	}
	return u
}

// Format creates the complete URL
func (u *URL) Format() string {
	_u, err := url.Parse(u.URL)
	if err != nil {
		log.Fatalf("%s", Pretty(err))
	}

	q := _u.Query()
	for key, value := range u.GetParameters {
		q.Set(key, fmt.Sprintf("%v", value))
	}

	_u.RawQuery = q.Encode()

	return _u.String()
}

// urlJoinRemoveExtremeSlashes removes first and last slash of a URL part
// useful for joining chunks to make a complete URL
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
