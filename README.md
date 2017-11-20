# Frank

[![Build Status](https://travis-ci.org/txthinking/frank.svg?branch=master)](https://travis-ci.org/txthinking/frank) [![Go Report Card](https://goreportcard.com/badge/github.com/txthinking/frank)](https://goreportcard.com/report/github.com/txthinking/frank) [![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](http://www.gnu.org/licenses/gpl-3.0) [![Wiki](https://img.shields.io/badge/docs-wiki-blue.svg)](https://github.com/txthinking/frank/wiki)

### Table of Contents

* [What is Frank](#what-is-frank)
* [Install](#install)
* [Test case file](#test-case-file)
	* [Score](#score)
	* [Comment](#comment)
	* [Functions](functions)
* [Example](#full-example)
	* [Simple example](#simple-example)
	* [POST form](#post-form)
	* [POST form file](#post-form-file)
	* [POST json data](#post-json-data)
	* [Use variable in path](#use-variable-in-path)
	* [Parse JSON body](#parse-json-body)
	* [Create variable and use it later](#create-variable-and-use-it-later)
	* [Use builtin function](#use-builtin-function)
	* [Print](#print)
* [Run case](#run-case)
* [Generate markdown document](#generate-markdown-document)
* [Contributing](#contributing)
* [License](#license)

## What is Frank

Frank is REST API automated testing tool like Postman but in command line.

## Install

```
$ sudo curl -L git.io/frank -o /usr/local/bin/frank
$ sudo chmod +x /usr/local/bin/frank
```

## Test case file

Test case file is actually a special javascript file.

### Score

It contains there scope:

* Init Score
	* `Init Score` can define some variables if needed before request started.
	* Must define `url` variable.
* Request Score
	* `Request Score` starts with a line that begines with `GET `, `HEAD `, `OPTIONS `, `POST `, `PUT ` or `DELETE `
	* The start line format `METHOD PATH [NAME]`, name is optionnal
		*  like like this: `GET /path` or `GET /path Name this request`
		* `PATH` must not have `?`, query and fragment
	* This score contain this predefined variables
		* `header={}` used for http header
		* `bounday={}` used for `header["Content-Type"] = "multipart/form-data; boundary=" + boundary`
		* `query={}` used for http query parameters
		* `form={}` used for http body when `Content-Type` is `application/x-www-form-urlencoded` or `multipart/form-data`
			* If form contain file, file name must start with `@`, like this: `form.key="@/path/to/file"`
		* `json={}` used for http body when `Content-Type` is `application/json`
		* `bodyRaw=""` a string used for http body, if this is not empty then use it and ignore `form`, `json` and `bodyFile`
		* `bodyFile=""` a file name, contents of file used for http body, if this is not empty then use it and ignore `form`, `json` and `bodyRaw`
	* This variables will be reset when Request Score starts
	* Request Score must be in pairs with Response Score
* Response Score
	* `Response Score` starts with a line that begines with `Response`
	* This score contain this predefined variables
		* `status` int, http status code
		* `proto` string, http protocol, like `HTTP/2.0`
		* `header` object, http header
		* `body` string, http body
	* This variables will be reassign When response Score starts
	* Response Score must be in pairs with Request Score

### Comment

You can use `//` to comment a line, like this:
```
// This is a comment line
```

### Functions

* `md5` Arguments: string. Return: string.
* `exit` Arguments: No. Return: No.
* `must` Arguments: boolean. Return: No. If pass argument is not equal to `true`, frank will exit immediately.

## Example

### Simple example

```
url="https://httpbin.org"

GET /ip

Response
```

### POST form

```
url="https://httpbin.org"

POST /post Post form
header["Content-Type"] = "application/x-www-form-urlencoded"
form.key0 = "value0"
form.key1 = "value1"

Response
must(status==200)
```

### POST form file

```
url="https://httpbin.org"

POST /post Post file
header["Content-Type"] = "multipart/form-data; boundary=" + boundary
form.key0 = "value0"
form.key1 = "@/etc/hosts"

Response
must(status==200)
```

### POST json data

```
url="https://httpbin.org"

POST /post Post json data
header["Accept"] = "application/json"
header["Content-Type"] = "application/json"
json.key0 = "value0"
json.key1 = "value1"

Response
must(status==200)
```

### Use variable in path

```
url="https://httpbin.org"
some="thing"

GET /etag/:some Just a GET request

Response
must(status==200)
```

### Parse JSON

```
url="https://httpbin.org"
some="thing"

POST /post Post json data
header["Accept"] = "application/json"
header["Content-Type"] = "application/json"
json.key0 = "value0"
json.key1 = some

Response
must(status==200)
j = JSON.parse(body)
must(j.origin.length > 3)
```

### Create variable and use it later

```
url="https://httpbin.org"

POST /post Post json data
header["Accept"] = "application/json"
header["Content-Type"] = "application/json"
json.key0 = "value0"
json.key1 = "value1"

Response
must(status==200)
j = JSON.parse(body)
myip = j.origin // created a new variable

GET /etag/:myip Just a GET request
query.key0 = myip // use a varible you created earlier

Response
must(status==200)
```

### Use builtin function

```
url="https://httpbin.org"

GET /ip
query.key0 = md5("value0")

Response
must(status==200)
```

### Print

```
url="https://httpbin.org"

GET /ip

Response
must(status==200)
console.log(body)
```

## Run case
```
# Default case file is ./test.frank
$ frank

# Specifies case file path
$ frank -c /path/to/case/file

# Set 500ms request interval
$ frank -d 500
```

## Generate markdown document
```
# Print to stdout
$ frank -m

# Write to api.md
$ frank -m > api.md
```

## Contributing

* Please create PR on `develop` branch

## License

Licensed under The GPLv3 License
