package goutil

import (
	"net/http"
)

// FindURLAfterRedirection gets an URL and returns the final URL after redirection
func FindURLAfterRedirection(url string) (*string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// The Request in the Response is the last URL the client tried to access.
	finalURL := resp.Request.URL.String()

	return &finalURL, nil
}
