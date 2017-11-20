package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	sj "github.com/bitly/go-simplejson"
)

// Response is one of responses in case
type Response struct {
	StatusCode int
	Proto      string
	Header     http.Header
	Cookies    []*http.Cookie
	Body       []byte
	JS         string
}

// ResponseBegin determines whether the response begined
func ResponseBegin(s string) bool {
	if strings.HasPrefix(s, "Response") {
		return true
	}
	return false
}

// CopyFrom copies from a http response
func (r *Response) CopyFrom(res *http.Response) error {
	defer res.Body.Close()
	r.StatusCode = res.StatusCode
	r.Proto = res.Proto
	r.Header = res.Header
	r.Cookies = res.Cookies()
	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.Header.Get("Content-Type") == "application/json" {
		j, err := sj.NewJson(d)
		if err != nil {
			return err
		}
		d, err = j.EncodePretty()
		if err != nil {
			return err
		}
	}
	r.Body = d
	if !markdown {
		PrintResponse(res, r)
	} else {
		PrintResponseMarkdown(res, r)
	}
	return nil
}

// Parse parses response
func (r *Response) Parse() error {
	if err := VM.Set("status", r.StatusCode); err != nil {
		return err
	}
	if err := VM.Set("proto", r.Proto); err != nil {
		return err
	}
	o, err := VM.Object(`({})`)
	if err != nil {
		return err
	}
	for k, v := range r.Header {
		o.Set(k, v[0])
	}
	if err := VM.Set("header", o); err != nil {
		return err
	}
	o, err = VM.Object(`({})`)
	if err != nil {
		return err
	}
	for _, v := range r.Cookies {
		o.Set(v.Name, v.Value)
	}
	if err := VM.Set("cookie", o); err != nil {
		return err
	}
	if err := VM.Set("body", string(r.Body)); err != nil {
		return err
	}
	if _, err := VM.Run(r.JS); err != nil {
		return err
	}
	return nil
}
