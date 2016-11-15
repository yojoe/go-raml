Benchmark Results

[rakyll/hey]() is used as benchmarking tool.

See the benchmark results below.

## Plaintext

The server endpoint simply return `Hello World!` string.

hey command : `hey -n 30000 -c 50 http://localhost:5000` which means:

- GET request
- number of request : 30K
- number of concurrent request : 50

**Go 4 CPUs**

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

**Go 1 CPU **

```
Summary:
  Total:    2.2095 secs
  Slowest:    0.0127 secs
  Fastest:  0.0001 secs
  Average:    0.0036 secs
  Requests/sec: 13577.7932
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

**Python**
```
Summary:
  Total:  247.2652 secs
  Slowest:  13.6994 secs
  Fastest:  0.0016 secs
  Average:  0.3942 secs
  Requests/sec: 121.3272
  Total data: 360000 bytes
  Size/request: 12 bytes
```

## Returns JSON

The server endpoint returns JSON array

hey command : `hey -n 30000 -c 50 http://localhost:5000/users`

**Go 4 CPU**
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

**Go 1 CPU**

```
Summary:
  Total:  2.4606 secs
  Slowest:  0.0136 secs
  Fastest:  0.0001 secs
  Average:  0.0040 secs
  Requests/sec: 12192.2035
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

**Python**

```
Summary:
  Total:  243.2155 secs
  Slowest:  13.8744 secs
  Fastest:  0.0017 secs
  Average:  0.3826 secs
  Requests/sec: 123.3474
  Total data: 1260000 bytes
  Size/request: 42 bytes
```

## Decode JSON

The endpoint decode a JSON string and response with plain string

command : `hey -n 30000 -c 50 -m POST -d '{"name":"John", "username": "Doe"}'  http://localhost:5000/users`

**Go 4 CPUs**

```
Summary:
  Total:  3.2592 secs
  Slowest:  0.0375 secs
  Fastest:  0.0002 secs
  Average:  0.0053 secs
  Requests/sec: 9204.7225
  Total data: 210000 bytes
  Size/request: 7 bytes
```

**Go 1 CPU**

```
Summary:
  Total:  3.3707 secs
  Slowest:  0.0324 secs
  Fastest:  0.0002 secs
  Average:  0.0055 secs
  Requests/sec: 8900.3395
  Total data: 210000 bytes
  Size/request: 7 bytes
```

**Nim**

Hey use chunked transfer encoding which not supported by Nim HTTP library

**Python**

```
Summary:
  Total:  242.1206 secs
  Slowest:  7.3893 secs
  Fastest:  0.0020 secs
  Average:  0.3891 secs
  Requests/sec: 123.9052
```