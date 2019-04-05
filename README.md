# Frank

[![Build Status](https://travis-ci.org/txthinking/frank.svg?branch=master)](https://travis-ci.org/txthinking/frank) [![Go Report Card](https://goreportcard.com/badge/github.com/txthinking/frank)](https://goreportcard.com/report/github.com/txthinking/frank) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](http://www.gnu.org/licenses/gpl-3.0)

### Table of Contents

* [What is Frank](#what-is-frank)
* [Install](#install)
	* [Linux](#linux)
	* [MacOS](#macos)
	* [Source](#source)
* [Test case file](#test-case-file)
	* [Score](#score)
	* [Comment](#comment)
	* [Builtin functions](#builtin-functions)
* [Example](#example)
	* [Simple example](#simple-example)
	* [POST form](#post-form)
	* [POST form file](#post-form-file)
	* [POST json data](#post-json-data)
	* [Use variable in path](#use-variable-in-path)
	* [Parse JSON](#parse-json)
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

### Linux

```
$ sudo curl -L git.io/frank -o /usr/local/bin/frank
$ sudo chmod +x /usr/local/bin/frank
```

### MacOS

```
$ curl -L git.io/frank_macos -o /usr/local/bin/frank
$ chmod +x /usr/local/bin/frank
```

### Source

```
$ go get github.com/txthinking/frank
```

## Test case file

Test case file is actually a special javascript file.

### Score

Frank contains there scopes:

#### Init Score

`Init Score` can define some variables if needed before request started.

* Must define `url` variable.

#### Request Score

The start line format `METHOD PATH [NAME]`, name is optionnal

* `Request Score` starts with a line that begins with `GET `, `HEAD `, `OPTIONS `, `POST `, `PUT `, `PATCH` or `DELETE `
* `PATH` must have no `?`, query and fragment

Predefined variables

| Name | Type | Description |
| --- | --- | --- |
| `header` | `object` | used for http header |
| `bounday` | `string` | used for `header["Content-Type"] = "multipart/form-data; boundary=" + boundary` |
| `query` | `object` | used for http query parameters |
| `form` | `object` | used for http body when `Content-Type` is `application/x-www-form-urlencoded` or `multipart/form-data` |
| `json` | `object` | used for http body when `Content-Type` is `application/json` |
| `bodyRaw` | `string` | used for http body, if this is not empty then use it and ignore `form`, `json` and `bodyFile` |
| `bodyFile` | `string` | a file path, contents of file used for http body, if this is not empty then use it and ignore `form`, `json` and `bodyRaw` |

> This variables will be reset when `Request Score` starts<br/>
> `Request Score` must be in pairs with `Response Score`

#### Response Score

`Response Score` starts with a line that begins with `Response`

Predefined variables

| Name | Type | Description |
| --- | --- | --- |
| `status` | `int` | http status code |
| `proto` | `string` | http protocol, like `HTTP/2.0` |
| `header` | `object` | http header |
| `cookie` | `object` | http cookies key/value |
| `body` | `string` | http body |

> This variables will be reassigned when `Response Score` starts<br/>
> `Response Score` must be in pairs with `Request Score`

### Comment

```
// This is a comment line
```

### Builtin functions

| Name | Arguments | Return value | Description |
| --- | --- | --- | --- |
| `base64encode` | string |  string | standard base64 encode |
| `base64decode` | string |  string | standard base64 decode |
| `exit` | - | - | Exit immediately with code 0 |
| `md5` | string |  string | md5 hash |
| `must` | boolean | - | If argument is not equal to `true`, will exit immediately with code 2 |

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
list=[]
list.push("value3")
list.push("value4")
json.key2 = list
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

GET /get
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

Please create PR on `develop` branch

## License

Licensed under The MIT License
