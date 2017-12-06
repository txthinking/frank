package main

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

// HttpClient is a http client
type HttpClient struct {
	Client *http.Client
	URL    string
}

// NewHttpClient return a new client
func NewHttpClient(url string) *HttpClient {
	c := &HttpClient{
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
		URL: url,
	}
	return c
}

// Do a request
func (c *HttpClient) Do(req *Request) (*http.Response, error) {
	var body io.Reader
	if req.Method == "POST" || req.Method == "PUT" || req.Method == "PATCH" {
		body = bytes.NewReader(req.Body)
	}
	q := "?" + req.Query.Encode()
	if q == "?" {
		q = ""
	}
	r, err := http.NewRequest(req.Method, c.URL+req.Path+q, body)
	if err != nil {
		return nil, err
	}
	r.Header = req.Header
	res, err := c.Client.Do(r)
	if err != nil {
		return nil, err
	}
	if !markdown {
		PrintRequest(r, req)
	} else {
		PrintRequestMarkdown(r, req)
	}
	return res, nil
}
