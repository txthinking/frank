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
Access-Control-Allow-Origin: *
X-Powered-By: Flask
Content-Length: 242
Via: 1.1 vegur
Server: meinheld/0.6.1
Content-Type: application/json
Access-Control-Allow-Credentials: true
X-Processed-Time: 0.00106287002563
Connection: keep-alive
Date: Mon, 20 Nov 2017 14:20:05 GMT

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
User-Agent: github.com/txthinking/frank
X-Hello-Frank: Frank

```
</details>
<details>
<summary>Response</summary>

```
HTTP/1.1 200 OK
Connection: keep-alive
Server: meinheld/0.6.1
Access-Control-Allow-Credentials: true
Content-Length: 33
Via: 1.1 vegur
Date: Mon, 20 Nov 2017 14:20:06 GMT
Content-Type: application/json
Access-Control-Allow-Origin: *
X-Powered-By: Flask
X-Processed-Time: 0.000530004501343

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
Server: meinheld/0.6.1
Date: Mon, 20 Nov 2017 14:20:06 GMT
X-Powered-By: Flask
Via: 1.1 vegur
Content-Type: application/json
Access-Control-Allow-Origin: *
Access-Control-Allow-Credentials: true
X-Processed-Time: 0.000771999359131
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
Content-Type: multipart/form-data; boundary=3881164518157487163

--3881164518157487163
Content-Disposition: form-data; name="key0"

valueo
--3881164518157487163
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

--3881164518157487163--

```
</details>
<details>
<summary>Response</summary>

```
HTTP/1.1 200 OK
X-Processed-Time: 0.00169610977173
Content-Type: application/json
Access-Control-Allow-Credentials: true
X-Powered-By: Flask
Access-Control-Allow-Origin: *
Content-Length: 733
Via: 1.1 vegur
Connection: keep-alive
Server: meinheld/0.6.1
Date: Mon, 20 Nov 2017 14:20:06 GMT

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
    "Content-Type": "multipart/form-data; boundary=3881164518157487163",
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
Connection: keep-alive
Date: Mon, 20 Nov 2017 14:20:05 GMT
Access-Control-Allow-Origin: *
Access-Control-Allow-Credentials: true
X-Powered-By: Flask
Via: 1.1 vegur
Server: meinheld/0.6.1
Content-Type: application/json
X-Processed-Time: 0.00124096870422
Content-Length: 555

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
Date: Mon, 20 Nov 2017 14:20:06 GMT
Etag: 104.199.139.23
Access-Control-Allow-Origin: *
X-Processed-Time: 0.00149393081665
Via: 1.1 vegur
Connection: keep-alive
Server: meinheld/0.6.1
Content-Type: application/json
Access-Control-Allow-Credentials: true
X-Powered-By: Flask
Content-Length: 310

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

<br/>

> Generated by [https://github.com/txthinking/frank](https://github.com/txthinking/frank)
