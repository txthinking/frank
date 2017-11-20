```
// This is a frank test case file
// -- https://github.com/txthinking/frank

url="https://httpbin.org"
some="thing"

```

### `GET /get`
<details>
<summary>Request</summary>

```
GET /get HTTP/1.1
User-Agent: github.com/txthinking/frank

```
</details>
<details>
<summary>Response</summary>

```
HTTP/1.1 200 OK
Connection: keep-alive
X-Powered-By: Flask
Content-Length: 242
Via: 1.1 vegur
Server: meinheld/0.6.1
Date: Mon, 20 Nov 2017 11:53:04 GMT
Content-Type: application/json
Access-Control-Allow-Origin: *
Access-Control-Allow-Credentials: true
X-Processed-Time: 0.000763177871704

{
  "args": {},
  "headers": {
    "Accept-Encoding": "gzip",
    "Connection": "close",
    "Host": "httpbin.org",
    "User-Agent": "github.com/txthinking/frank"
  },
  "origin": "104.199.139.23",
  "url": "https://httpbin.org/get"
}
```
</details>

### `GET /ip Get my IP`
<details>
<summary>Request</summary>

```
GET /ip?key0=value0 HTTP/1.1
X-Hello-Frank: Frank
User-Agent: github.com/txthinking/frank

```
</details>
<details>
<summary>Response</summary>

```
HTTP/1.1 200 OK
X-Processed-Time: 0.000442981719971
Via: 1.1 vegur
Date: Mon, 20 Nov 2017 11:53:04 GMT
Access-Control-Allow-Origin: *
Content-Type: application/json
Access-Control-Allow-Credentials: true
X-Powered-By: Flask
Content-Length: 33
Connection: keep-alive
Server: meinheld/0.6.1

{
  "origin": "104.199.139.23"
}
```
</details>

### `POST /post Post form`
<details>
<summary>Request</summary>

```
POST /post HTTP/1.1
User-Agent: github.com/txthinking/frank
Content-Type: application/x-www-form-urlencoded

key0=value0
```
</details>
<details>
<summary>Response</summary>

```
HTTP/1.1 200 OK
Connection: keep-alive
Date: Mon, 20 Nov 2017 11:53:04 GMT
Access-Control-Allow-Origin: *
X-Powered-By: Flask
X-Processed-Time: 0.000938177108765
Via: 1.1 vegur
Server: meinheld/0.6.1
Content-Type: application/json
Access-Control-Allow-Credentials: true
Content-Length: 417

{
  "args": {},
  "data": "",
  "files": {},
  "form": {
    "key0": "value0"
  },
  "headers": {
    "Accept-Encoding": "gzip",
    "Connection": "close",
    "Content-Length": "11",
    "Content-Type": "application/x-www-form-urlencoded",
    "Host": "httpbin.org",
    "User-Agent": "github.com/txthinking/frank"
  },
  "json": null,
  "origin": "104.199.139.23",
  "url": "https://httpbin.org/post"
}
```
</details>

### `POST /post Post file`
<details>
<summary>Request</summary>

```
POST /post HTTP/1.1
User-Agent: github.com/txthinking/frank
Content-Type: multipart/form-data; boundary=4675639043856151899

--4675639043856151899
Content-Disposition: form-data; name="key0"

valueo
--4675639043856151899
Content-Disposition: form-data; name="key1"; filename="hosts"
Content-Type: application/octet-stream

127.0.0.1 localhost

# The following lines are desirable for IPv6 capable hosts
::1 ip6-localhost ip6-loopback
fe00::0 ip6-localnet
ff00::0 ip6-mcastprefix
ff02::1 ip6-allnodes
ff02::2 ip6-allrouters
ff02::3 ip6-allhosts
169.254.169.254 metadata.google.internal metadata

--4675639043856151899--

```
</details>
<details>
<summary>Response</summary>

```
HTTP/1.1 200 OK
Access-Control-Allow-Credentials: true
Content-Length: 733
Connection: keep-alive
Date: Mon, 20 Nov 2017 11:53:04 GMT
Content-Type: application/json
Access-Control-Allow-Origin: *
X-Powered-By: Flask
X-Processed-Time: 0.00134897232056
Via: 1.1 vegur
Server: meinheld/0.6.1

{
  "args": {},
  "data": "",
  "files": {
    "key1": "127.0.0.1 localhost\n\n# The following lines are desirable for IPv6 capable hosts\n::1 ip6-localhost ip6-loopback\nfe00::0 ip6-localnet\nff00::0 ip6-mcastprefix\nff02::1 ip6-allnodes\nff02::2 ip6-allrouters\nff02::3 ip6-allhosts\n169.254.169.254 metadata.google.internal metadata\n"
  },
  "form": {
    "key0": "valueo"
  },
  "headers": {
    "Accept-Encoding": "gzip",
    "Connection": "close",
    "Content-Length": "504",
    "Content-Type": "multipart/form-data; boundary=4675639043856151899",
    "Host": "httpbin.org",
    "User-Agent": "github.com/txthinking/frank"
  },
  "json": null,
  "origin": "104.199.139.23",
  "url": "https://httpbin.org/post"
}
```
</details>

### `POST /post Post json data`
<details>
<summary>Request</summary>

```
POST /post HTTP/1.1
User-Agent: github.com/txthinking/frank
Accept: application/json
Content-Type: application/json

{
  "key0": "value0",
  "key1": "value1",
  "key2": "thing"
}
```
</details>
<details>
<summary>Response</summary>

```
HTTP/1.1 200 OK
Server: meinheld/0.6.1
Access-Control-Allow-Origin: *
Access-Control-Allow-Credentials: true
X-Powered-By: Flask
Content-Length: 555
Connection: keep-alive
Date: Mon, 20 Nov 2017 11:53:05 GMT
Content-Type: application/json
X-Processed-Time: 0.00137686729431
Via: 1.1 vegur

{
  "args": {},
  "data": "{\n  \"key0\": \"value0\",\n  \"key1\": \"value1\",\n  \"key2\": \"thing\"\n}",
  "files": {},
  "form": {},
  "headers": {
    "Accept": "application/json",
    "Accept-Encoding": "gzip",
    "Connection": "close",
    "Content-Length": "61",
    "Content-Type": "application/json",
    "Host": "httpbin.org",
    "User-Agent": "github.com/txthinking/frank"
  },
  "json": {
    "key0": "value0",
    "key1": "value1",
    "key2": "thing"
  },
  "origin": "104.199.139.23",
  "url": "https://httpbin.org/post"
}
```
</details>

### `GET /etag/:myip Just a GET request`
<details>
<summary>Request</summary>

```
GET /etag/104.199.139.23?key0=104.199.139.23 HTTP/1.1
User-Agent: github.com/txthinking/frank

```
</details>
<details>
<summary>Response</summary>

```
HTTP/1.1 200 OK
Date: Mon, 20 Nov 2017 11:53:05 GMT
Access-Control-Allow-Origin: *
Access-Control-Allow-Credentials: true
X-Powered-By: Flask
Content-Length: 310
Via: 1.1 vegur
Connection: keep-alive
Server: meinheld/0.6.1
Content-Type: application/json
Etag: 104.199.139.23
X-Processed-Time: 0.000746965408325

{
  "args": {
    "key0": "104.199.139.23"
  },
  "headers": {
    "Accept-Encoding": "gzip",
    "Connection": "close",
    "Host": "httpbin.org",
    "User-Agent": "github.com/txthinking/frank"
  },
  "origin": "104.199.139.23",
  "url": "https://httpbin.org/etag/104.199.139.23?key0=104.199.139.23"
}
```
</details>

> -- Generated by https://github.com/txthinking/frank
