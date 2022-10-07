package ixapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// bodyReader helper
type bodyReader struct {
	*bytes.Reader
}

// Close implements the reader closer interface
func (r *bodyReader) Close() error {
	return nil
}

// Create body reader from bytes
func newBodyReader(body []byte) *bodyReader {
	r := bytes.NewReader(body)
	return &bodyReader{Reader: r}
}

// TestResponseFunc is a test response generator
type TestResponseFunc func(body []byte) (any, error)

// NewTestClient creates an IX-API client
// with a custom response for a request endpoint.
// The response will be encoded as json.
func NewTestClient(responses map[string]any) *Client {
	c := NewClient("") // plain http client
	c.Transport = NewMockResponseTransport(responses)
	return c
}

// MockResponseTransport implements the RoundTripper interface
// and will create a response for a request to a specific endpoint
type MockResponseTransport struct {
	responses map[string]any
}

// NewMockResponseTransport creates a transport which
// will return mocked response data
func NewMockResponseTransport(responses map[string]any) *MockResponseTransport {
	return &MockResponseTransport{
		responses: responses,
	}
}

// RoundTrip implements the transports RoundTripper interface
// and creates responses for requests
func (t *MockResponseTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Println(req)
	var data any
	resHandler, ok := t.responses[req.URL.Path]
	if !ok {
		return NewNotFoundResponse(req), nil
	}

	// Response can be a function
	handlerFunc, ok := resHandler.(TestResponseFunc)
	if ok {
		var body []byte
		if req.Body != nil {
			body, _ = ioutil.ReadAll(req.Body)
		}
		d, err := handlerFunc(body)
		if err != nil {
			return nil, err
		}
		data = d
	} else {
		data = resHandler
	}

	body, _ := json.Marshal(data)

	res := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Body:          newBodyReader(body),
		ContentLength: int64(len(body)),
	}

	return res, nil
}

// NewNotFoundResponse creates a new 404 not found response with body
func NewNotFoundResponse(req *http.Request) *http.Response {
	err := &ProblemResponse{
		Type:     "not_found",
		Title:    "The requested resource was not found",
		Status:   404,
		Detail:   "Could not find the requested resource",
		Instance: req.URL.String(),
	}
	errorBody, _ := json.Marshal(err)
	res := &http.Response{
		Status:        "404 NotFound",
		StatusCode:    404,
		Body:          newBodyReader(errorBody),
		ContentLength: int64(len(errorBody)),
	}
	return res
}
