package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Request struct {
	RawName string
	Name    string
	Method  string
	Path    string
	Header  http.Header
	Query   url.Values
	Body    []byte
	JS      string
}

type Response struct {
	StatusCode int
	Proto      string
	Header     http.Header
	Body       []byte
	JS         string
}

type RequestResponse struct {
	Request  *Request
	Response *Response
}

type Case struct {
	GlobalJS string
	RR       []*RequestResponse
	Client   *HttpClient
}

func NewCase(path string) (*Case, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	bf := bytes.NewBuffer(b)
	js := ""
	var req *Request
	var res *Response
	c := &Case{}
	c.RR = make([]*RequestResponse, 0)
	for {
		l, err := bf.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			res = &Response{}
			res.JS = js
			c.RR = append(c.RR, &RequestResponse{
				Request:  req,
				Response: res,
			})
			return c, nil
		}
		if RequestBegin(l) {
			if c.GlobalJS == "" {
				c.GlobalJS = js
			} else {
				res = &Response{}
				res.JS = js
				c.RR = append(c.RR, &RequestResponse{
					Request:  req,
					Response: res,
				})
			}
			js = l
			continue
		}
		if ResponseBegin(l) {
			req = &Request{}
			req.JS = js
			js = ""
			continue
		}
		js += l
	}
}

func (c *Case) Run() error {
	e := func(r *Request, err error) error {
		return errors.New(fmt.Sprintf("%s on %s request", err.Error(), r.Name))
	}
	if markdown {
		fmt.Printf("```\n%s```\n", c.GlobalJS)
	}
	_, err := VM.Run(c.GlobalJS)
	if err != nil {
		return err
	}
	v, err := VM.Get("url")
	if err != nil {
		return err
	}
	if !v.IsString() {
		return errors.New("Invalid url")
	}
	url, err := v.ToString()
	if err != nil {
		return err
	}
	c.Client = NewHttpClient(url)
	for _, v := range c.RR {
		if err := v.Request.MakeStartLine(); err != nil {
			return e(err)
		}
		if err := v.Request.Parse(); err != nil {
			return e(err)
		}
		if err := v.Request.MakeHeader(); err != nil {
			return e(err)
		}
		if err := v.Request.MakeQuery(); err != nil {
			return e(err)
		}
		if v.Request.Method == "POST" || v.Request.Method == "PUT" {
			if err := v.Request.MakeBody(); err != nil {
				return e(err)
			}
		}

		res, err := c.Client.Do(v.Request)
		if err != nil {
			return e(err)
		}
		if err := v.Response.CopyFrom(res); err != nil {
			return e(err)
		}
		if err := v.Response.Parse(); err != nil {
			return e(err)
		}
	}
	return nil
}
