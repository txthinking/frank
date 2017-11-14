package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	sj "github.com/bitly/go-simplejson"
	"github.com/fatih/color"
)

func RequestBegin(s string) bool {
	if !strings.HasPrefix(s, "DELETE ") &&
		!strings.HasPrefix(s, "GET ") &&
		!strings.HasPrefix(s, "HEAD ") &&
		!strings.HasPrefix(s, "OPTIONS ") &&
		!strings.HasPrefix(s, "POST ") &&
		!strings.HasPrefix(s, "PUT ") {
		return false
	}
	return true
}

func ResponseBegin(s string) bool {
	if strings.HasPrefix(s, "Response") {
		return true
	}
	return false
}

func (r *Request) MakeStartLine() error {
	ss := strings.SplitN(r.JS, "\n", 2)
	if len(ss) != 2 {
		return errors.New("Invalid format")
	}
	r.RawName = ss[0]
	r.JS = ss[1]
	ss = strings.SplitN(ss[0], " ", 3)
	if len(ss) < 2 {
		return errors.New("Invalid format")
	}
	r.Method = ss[0]
	if len(ss) == 3 {
		r.Name = ss[2]
	}

	re, err := regexp.Compile(`:[A-Za-z_]\w*`)
	if err != nil {
		return err
	}
	p := ss[1]
	ss = re.FindAllString(ss[1], -1)
	for _, v := range ss {
		vl, err := VM.Get(v[1:])
		if err != nil {
			return err
		}
		va, err := vl.ToString()
		if err != nil {
			return err
		}
		p = strings.Replace(p, v, va, -1)
	}
	r.Path = p
	return nil
}

func (r *Request) Parse() error {
	_, err := VM.Run(`header = {"User-Agent": "github.com/txthinking/frank"}; query = {}; param = {};`)
	if err != nil {
		return err
	}
	_, err = VM.Run(r.JS)
	if err != nil {
		return err
	}
	return nil
}

func (r *Request) MakeHeader() error {
	tmp, err := VM.Get("header")
	if err != nil {
		return err
	}
	if !tmp.IsObject() {
		return err
	}
	a, err := tmp.Export()
	if err != nil {
		return err
	}
	m, ok := a.(map[string]interface{})
	if !ok {
		return errors.New("Invalid header")
	}
	h := http.Header{}
	for k, v := range m {
		s, ok := v.(string)
		if !ok {
			i, ok := v.(int64)
			if !ok {
				return errors.New("Invalid header")
			}
			s = strconv.Itoa(int(i))
		}
		h.Set(http.CanonicalHeaderKey(k), s)
	}
	r.Header = h
	return nil
}

func (r *Request) MakeQuery() error {
	tmp, err := VM.Get("query")
	if err != nil {
		return err
	}
	if !tmp.IsObject() {
		return err
	}
	a, err := tmp.Export()
	if err != nil {
		return err
	}
	m, ok := a.(map[string]interface{})
	if !ok {
		return errors.New("Invalid query")
	}
	vl := url.Values{}
	for k, v := range m {
		s, ok := v.(string)
		if !ok {
			i, ok := v.(int64)
			if !ok {
				return errors.New("Invalid query")
			}
			s = strconv.Itoa(int(i))
		}
		vl.Set(k, s)
	}
	r.Query = vl
	return nil
}

func (r *Request) MakeBody() error {
	tmp, err := VM.Get("param")
	if err != nil {
		return err
	}
	if !tmp.IsObject() {
		return err
	}
	a, err := tmp.Export()
	if err != nil {
		return err
	}
	m, ok := a.(map[string]interface{})
	if !ok {
		return errors.New("Invalid query")
	}
	ct := r.Header.Get("Content-Type")
	if ct == "application/json" {
		d, err := json.Marshal(m)
		if err != nil {
			return err
		}
		j, err := sj.NewJson(d)
		if err != nil {
			return err
		}
		d, err = j.EncodePretty()
		if err != nil {
			return err
		}
		r.Body = d
		return nil
	}
	if ct == "application/x-www-form-urlencoded" {
		vl := url.Values{}
		for k, v := range m {
			s, ok := v.(string)
			if !ok {
				i, ok := v.(int64)
				if !ok {
					return errors.New("Invalid param")
				}
				s = strconv.Itoa(int(i))
			}
			vl.Set(k, s)
		}
		r.Body = []byte(vl.Encode())
		return nil
	}
	return errors.New("Unsupport Content-Type")
}

func (r *Response) CopyFrom(res *http.Response) error {
	r.StatusCode = res.StatusCode
	r.Proto = res.Proto
	r.Header = res.Header
	defer res.Body.Close()
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
	if err := VM.Set("body", string(r.Body)); err != nil {
		return err
	}
	if _, err := VM.Run(r.JS); err != nil {
		return err
	}
	return nil
}

func PrintResponse(r *http.Response, res *Response) {
	color.Green("Response>")
	fmt.Printf("%s %s\r\n", r.Proto, r.Status)
	for k, _ := range r.Header {
		fmt.Printf("%s: %s\r\n", k, r.Header.Get(k))
	}
	fmt.Printf("\r\n")
	fmt.Printf("%s\n", res.Body)
}

func PrintResponseMarkdown(r *http.Response, res *Response) {
	fmt.Print("<details>\n")
	fmt.Print("<summary>Response</summary>\n")
	fmt.Print("\n")
	fmt.Print("```\n")
	fmt.Printf("%s %s\r\n", r.Proto, r.Status)
	for k, _ := range r.Header {
		fmt.Printf("%s: %s\r\n", k, r.Header.Get(k))
	}
	fmt.Printf("\r\n")
	fmt.Printf("%s\n", res.Body)
	fmt.Print("```\n")
	fmt.Print("</details>\n")
}
