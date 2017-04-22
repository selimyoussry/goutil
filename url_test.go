package goutil

import "testing"

func TestURL(t *testing.T) {

	urls := []string{}
	urls = append(urls, URLJoin("http://google.com/", "a", "b/", "c/d/e/", "f/g"))
	urls = append(urls, URLJoin("http://google.com", "a", "b/", "/c/d/e/", "f/g"))
	urls = append(urls, URLJoin("http://google.com", "a", "", "b/", "", "c/d/e/", "/f/g/"))

	expectedURL := "http://google.com/a/b/c/d/e/f/g"
	for i, url := range urls {
		if url != expectedURL {
			t.Errorf("Expected %s - Got %s - For #%d", expectedURL, url, i)
		}
	}
}
