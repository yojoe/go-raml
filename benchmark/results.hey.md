Benchmark Results

[rakyll/hey]() is used as benchmarking tool.

See the benchmark results below.

## Plaintext

The server endpoint simply return `Hello World!` string.

hey command : `hey -n 30000 -c 50 http://localhost:5000` which means:

- GET request
- number of request : 30K
- number of concurrent request : 50

**Go**

```
Summary:
  Total:    1.5014 secs
  Slowest:    0.0539 secs
  Fastest:  0.0001 secs
  Average:    0.0024 secs
  Requests/sec: 19981.8522
  Total data: 360000 bytes
  Size/request: 12 bytes
```

**Nim**

```
Summary:
  Total:    2.1525 secs
  Slowest:    0.0435 secs
  Fastest:  0.0002 secs
  Average:    0.0036 secs
  Requests/sec: 13937.1186
  Total data: 360000 bytes
  Size/request: 12 bytes
```

## Returns JSON

The server endpoint returns JSON array

hey command : `hey -n 30000 -c 50 http://localhost:5000/users`

**Go**
```
Summary:
  Total:    1.6139 secs
  Slowest:    0.0440 secs
  Fastest:  0.0001 secs
  Average:    0.0026 secs
  Requests/sec: 18588.8480
  Total data: 990000 bytes
  Size/request: 33 bytes
```

**Nim**

```
Summary:
  Total:    2.1811 secs
  Slowest:    0.0441 secs
  Fastest:  0.0001 secs
  Average:    0.0036 secs
  Requests/sec: 13754.6201
  Total data: 1050000 bytes
  Size/request: 35 bytes
```

## Decode JSON

The endpoint decode a JSON string and response with plain string

command : `hey -n 30000 -c 50 -m POST -d '{"name":"John", "username": "Doe"}'  http://localhost:5000/users`

**Go**

```
Summary:
  Total:    3.4145 secs
  Slowest:    0.0436 secs
  Fastest:  0.0002 secs
  Average:    0.0055 secs
  Requests/sec: 8786.1079
  Total data: 210000 bytes
  Size/request: 7 bytes
```

**Nim**
Failed to decode, with this error

Unsolicited response received on idle HTTP channel starting with "HTTP/1.1 400 Bad Request\r\nContent-Length: 31\r\n\r\nBad Request. No Content-Length."; err=<nil>

