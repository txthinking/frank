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
	"github.com/txthinking/x"
)

// Request is one of requests in case
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

// RequestBegin determines whether the request begined
func RequestBegin(s string) bool {
	if !strings.HasPrefix(s, "DELETE ") &&
		!strings.HasPrefix(s, "GET ") &&
		!strings.HasPrefix(s, "HEAD ") &&
		!strings.HasPrefix(s, "OPTIONS ") &&
		!strings.HasPrefix(s, "POST ") &&
		!strings.HasPrefix(s, "PATCH ") &&
		!strings.HasPrefix(s, "PUT ") {
		return false
	}
	return true
}

// MakeStartLine parses start-line
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

// Parse parses request
func (r *Request) Parse() error {
	_, err := VM.Run(fmt.Sprintf(`header={"User-Agent": "github.com/txthinking/frank"}; boundary="%d"; query={}; form={}; json={}; bodyRaw=""; bodyFile="";`, x.RandomNumber()))
	if err != nil {
		return err
	}
	_, err = VM.Run(r.JS)
	if err != nil {
		return err
	}
	return nil
}

// MakeHeader parses header
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
		s := ""
		switch t := v.(type) {
		case string:
			s = t
		case int64:
			s = strconv.FormatInt(t, 10)
		case float64:
			s = strconv.FormatFloat(t, 'f', -1, 64)
		default:
			return errors.New("Invalid header value of key: " + k)
		}
		h.Set(http.CanonicalHeaderKey(k), s)
	}
	r.Header = h
	return nil
}

// MakeQuery parses query parameters
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
		s := ""
		switch t := v.(type) {
		case string:
			s = t
		case int64:
			s = strconv.FormatInt(t, 10)
		case float64:
			s = strconv.FormatFloat(t, 'f', -1, 64)
		default:
			return errors.New("Invalid query value of key: " + k)
		}
		vl.Set(k, s)
	}
	r.Query = vl
	return nil
}

// MakeBody parses body
func (r *Request) MakeBody() error {
	v, err := VM.Get("bodyRaw")
	if err != nil {
		return err
	}
	if !v.IsString() {
		return errors.New("Invalid bodyRaw")
	}
	raw, err := v.ToString()
	if err != nil {
		return err
	}
	if raw != "" {
		r.Body = []byte(raw)
		return nil
	}

	v, err = VM.Get("bodyFile")
	if err != nil {
		return err
	}
	if !v.IsString() {
		return errors.New("Invalid bodyFile")
	}
	fn, err := v.ToString()
	if err != nil {
		return err
	}
	if fn != "" {
		b, err := ioutil.ReadFile(fn)
		if err != nil {
			return err
		}
		r.Body = b
		return nil
	}

	ct := r.Header.Get("Content-Type")
	if ct == "application/json" {
		tmp, err := VM.Get("json")
		if err != nil {
			return err
		}
		if !tmp.IsObject() {
			return err
		}
		it, err := tmp.Export()
		if err != nil {
			return err
		}
		m, ok := it.(map[string]interface{})
		if !ok {
			return errors.New("Invalid json")
		}

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
		tmp, err := VM.Get("form")
		if err != nil {
			return err
		}
		if !tmp.IsObject() {
			return err
		}
		it, err := tmp.Export()
		if err != nil {
			return err
		}
		m, ok := it.(map[string]interface{})
		if !ok {
			return errors.New("Invalid form")
		}
		vl := url.Values{}
		for k, v := range m {
			s := ""
			switch t := v.(type) {
			case string:
				s = t
			case int64:
				s = strconv.FormatInt(t, 10)
			case float64:
				s = strconv.FormatFloat(t, 'f', -1, 64)
			default:
				return errors.New("Invalid form value of key: " + k)
			}
			vl.Set(k, s)
		}
		r.Body = []byte(vl.Encode())
		return nil
	}
	if strings.HasPrefix(ct, "multipart/form-data") {
		v, err := VM.Get("boundary")
		if err != nil {
			return err
		}
		if !v.IsString() {
			return errors.New("Invalid boundary")
		}
		bd, err := v.ToString()
		if err != nil {
			return err
		}

		tmp, err := VM.Get("form")
		if err != nil {
			return err
		}
		if !tmp.IsObject() {
			return err
		}
		it, err := tmp.Export()
		if err != nil {
			return err
		}
		m, ok := it.(map[string]interface{})
		if !ok {
			return errors.New("Invalid form")
		}
		params := make(map[string][]string)
		files := make(map[string][]string)
		for k, v := range m {
			s := ""
			switch t := v.(type) {
			case string:
				s = t
			case int64:
				s = strconv.FormatInt(t, 10)
			case float64:
				s = strconv.FormatFloat(t, 'f', -1, 64)
			default:
				return errors.New("Invalid form value of key: " + k)
			}
			if !strings.HasPrefix(s, "@") {
				ss := []string{s}
				params[k] = ss
				continue
			}
			ss := []string{s[1:]}
			files[k] = ss
		}
		src, err := x.MultipartFormDataFromFile(params, files, bd)
		if err != nil {
			return err
		}
		b, err := ioutil.ReadAll(src)
		if err != nil {
			return err
		}
		r.Body = b
		return nil
	}
	return errors.New("Unsupport Content-Type")
}
