package goutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/hippoai/goerr"
)

const (
	DefaultHTTPTimeoutSeconds = 10
	ErrMsgHTTPRequest         = "ERR_HTTP_REQUEST"
)

// ErrHTTPRequest returns an error when the request didn't work
func ErrHTTPRequest(statusCode int, responseDataString string, response *http.Response) error {
	return goerr.New(ErrMsgHTTPRequest, map[string]interface{}{
		"statusCode":         statusCode,
		"responseDataString": responseDataString,
		"response":           response,
	})
}

// HTTPConnection
type HTTPConnection struct {
	Token             *string       `json:"token"`
	AuthorizationType *string       `json:"authorizationType"`
	TimeoutSeconds    time.Duration `json:"timeoutSeconds"`
}

// NewHTTPConnection instanciates
func NewHTTPConnection() *HTTPConnection {
	return &HTTPConnection{
		TimeoutSeconds: time.Duration(DefaultHTTPTimeoutSeconds),
	}
}

// SetAuthorization sets the token and type of Authorization header
func (hc *HTTPConnection) SetAuthorization(authorizationType, token string) *HTTPConnection {
	hc.Token = &token
	hc.AuthorizationType = &authorizationType
	return hc
}

// SetTimeoutSeconds overwrites the timeout for the connection
func (hc *HTTPConnection) SetTimeoutSeconds(t int) *HTTPConnection {
	hc.TimeoutSeconds = time.Duration(t)
	return hc
}

// SendRaw -
func (hc *HTTPConnection) SendRaw(url string, payload []byte, method string, isJSON bool) (*http.Response, []byte, error) {

	// Net client with a timeout
	netClient := &http.Client{
		Timeout: time.Second * hc.TimeoutSeconds,
	}

	// Create the http request
	req, err := http.NewRequest(method, url, strings.NewReader(string(payload)))
	if err != nil {
		return nil, nil, err
	}

	if isJSON {
		// Set JSON header
		req.Header.Set("Content-Type", "application/json")
	}

	// Set authorization header
	if hc.Token != nil {
		req.Header.Set("Authorization", fmt.Sprintf("%s %s", *hc.AuthorizationType, *hc.Token))
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Get the response
	resp, err := netClient.Do(req)

	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	// Parse the body
	respBodyBytesArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	// Make sure the response status is correct
	if resp.StatusCode != http.StatusOK {
		return nil, nil, ErrHTTPRequest(resp.StatusCode, string(respBodyBytesArray), resp)
	}

	return resp, respBodyBytesArray, nil
}

// Send a payload and get the response
func (hc *HTTPConnection) Send(url string, payload interface{}, method string) (*http.Response, []byte, error) {

	// Net client with a timeout
	netClient := &http.Client{
		Timeout: time.Second * hc.TimeoutSeconds,
	}

	// Marshal the body into a bytes array
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, err
	}

	// Create the http request
	var req *http.Request
	if payload == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(b))
	}

	if err != nil {
		return nil, nil, err
	}

	// Set JSON header
	req.Header.Set("Content-Type", "application/json")

	// Set authorization header
	if hc.Token != nil {
		req.Header.Set("Authorization", fmt.Sprintf("%s %s", *hc.AuthorizationType, *hc.Token))
	}

	// Get the response
	resp, err := netClient.Do(req)

	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	// Parse the body
	respBodyBytesArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	// Make sure the response status is correct
	if resp.StatusCode != http.StatusOK {
		return nil, nil, ErrHTTPRequest(resp.StatusCode, string(respBodyBytesArray), resp)
	}

	return resp, respBodyBytesArray, nil
}

// PostWithAuth - SendWithAuth with POST
func (hc *HTTPConnection) PostWithAuth(url string, payload interface{}) (*http.Response, []byte, error) {
	return hc.Send(url, payload, "POST")
}

// PutWithAuth - SendWithAuth with PUT
func (hc *HTTPConnection) PutWithAuth(url string, payload interface{}) (*http.Response, []byte, error) {
	return hc.Send(url, payload, "PUT")
}

// DeleteWithAuth - SendWithAuth with DELETE
func (hc *HTTPConnection) DeleteWithAuth(url string, payload interface{}) (*http.Response, []byte, error) {
	return hc.Send(url, payload, "DELETE")
}

// PatchWithAuth - SendWithAuth with PATCH
func (hc *HTTPConnection) PatchWithAuth(url string, payload interface{}) (*http.Response, []byte, error) {
	return hc.Send(url, payload, "PATCH")
}

// GetWithAuth - SendWithAuth with GET
func (hc *HTTPConnection) GetWithAuth(url string) (*http.Response, []byte, error) {
	return hc.Send(url, nil, "GET")
}
