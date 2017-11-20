package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

// PrintRequest prints the http request
func PrintRequest(r *http.Request, req *Request) {
	color.Green("Request %s>\n", req.Name)
	fmt.Printf("%s %s %s\n", r.Method, r.URL.RequestURI(), r.Proto)
	for k, _ := range r.Header {
		fmt.Printf("%s: %s\n", k, r.Header.Get(k))
	}
	fmt.Printf("\n")
	if r.Method == "POST" || r.Method == "PUT" {
		color.Cyan("%s\n", req.Body)
	}
}

// PrintRequestMarkdown prints the http request markdown format
func PrintRequestMarkdown(r *http.Request, req *Request) {
	fmt.Print("\n")
	fmt.Printf("### `%s`\n", req.RawName)
	fmt.Print("<details>\n")
	fmt.Print("<summary>Request</summary>\n")
	fmt.Print("\n")
	fmt.Print("```\n")
	fmt.Printf("%s %s %s\n", r.Method, r.URL.RequestURI(), r.Proto)
	for k, _ := range r.Header {
		fmt.Printf("%s: %s\n", k, r.Header.Get(k))
	}
	fmt.Printf("\n")
	if r.Method == "POST" || r.Method == "PUT" {
		fmt.Printf("%s\n", req.Body)
	}
	fmt.Print("```\n")
	fmt.Print("</details>\n")
}

// PrintResponse prints the http response
func PrintResponse(r *http.Response, res *Response) {
	color.Green("Response>")
	fmt.Printf("%s %s\n", r.Proto, r.Status)
	for k, _ := range r.Header {
		fmt.Printf("%s: %s\n", k, r.Header.Get(k))
	}
	fmt.Printf("\n")
	color.Cyan("%s\n", res.Body)
}

// PrintResponseMarkdown prints the http response markdown format
func PrintResponseMarkdown(r *http.Response, res *Response) {
	fmt.Print("<details>\n")
	fmt.Print("<summary>Response</summary>\n")
	fmt.Print("\n")
	fmt.Print("```\n")
	fmt.Printf("%s %s\n", r.Proto, r.Status)
	for k, _ := range r.Header {
		fmt.Printf("%s: %s\n", k, r.Header.Get(k))
	}
	fmt.Printf("\n")
	fmt.Printf("%s\n", res.Body)
	fmt.Print("```\n")
	fmt.Print("</details>\n")
}
