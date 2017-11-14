package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/fatih/color"
)

type HttpClient struct {
	Client *http.Client
	URL    string
}

func NewHttpClient(url string) *HttpClient {
	c := &HttpClient{
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
		URL: url,
	}
	return c
}

func (c *HttpClient) Do(req *Request) (*http.Response, error) {
	var body io.Reader
	if req.Method == "POST" || req.Method == "PUT" {
		body = bytes.NewReader(req.Body)
	}
	r, err := http.NewRequest(req.Method, c.URL+req.Path+"?"+req.Query.Encode(), body)
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

func PrintRequest(r *http.Request, req *Request) {
	fmt.Print("\n")
	color.Green("Request [%s]>\n", req.Name)
	fmt.Printf("GET %s HTTP/1.1\r\n", r.URL.RequestURI())
	for k, _ := range r.Header {
		fmt.Printf("%s: %s\r\n", k, r.Header.Get(k))
	}
	fmt.Printf("\r\n")
	if r.Method == "POST" || r.Method == "PUT" {
		fmt.Printf("%s\n", req.Body)
	}
}

func PrintRequestMarkdown(r *http.Request, req *Request) {
	fmt.Print("\n")
	fmt.Printf("### `%s`\n", req.RawName)
	fmt.Print("<details>\n")
	fmt.Print("<summary>Request</summary>\n")
	fmt.Print("\n")
	fmt.Print("```\n")
	for k, _ := range r.Header {
		fmt.Printf("%s: %s\r\n", k, r.Header.Get(k))
	}
	fmt.Printf("\r\n")
	if r.Method == "POST" || r.Method == "PUT" {
		fmt.Printf("%s\n", req.Body)
	}
	fmt.Print("```\n")
	fmt.Print("</details>\n")
}
