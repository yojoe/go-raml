Benchmark Results

[rakyll/hey]() is used as benchmarking tool.

See the benchmark results below.

## Plaintext

The server endpoint simply return `Hello World!` string.

hey command : `hey -n 30000 -c 50 http://localhost:5000` which means:

- GET request
- number of request : 30K
- number of concurrent request : 50

**Go 2 CPUs**

```
Summary:
  Total:        2.2703 secs
  Slowest:      0.0291 secs
  Fastest:      0.0002 secs
  Average:      0.0037 secs
  Requests/sec: 13214.2607
  Total data:   360000 bytes
  Size/request: 12 bytes
```

**Go 1 CPU**

```
Summary:
  Total:        3.3555 secs
  Slowest:      0.0291 secs
  Fastest:      0.0003 secs
  Average:      0.0055 secs
  Requests/sec: 8940.5224
  Total data:   360000 bytes
  Size/request: 12 bytes
```

**Nim**

```
Summary:
  Total:        2.8036 secs
  Slowest:      0.0250 secs
  Fastest:      0.0002 secs
  Average:      0.0046 secs
  Requests/sec: 10700.4852
  Total data:   1050000 bytes
  Size/request: 35 bytes
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

**Go 2 CPUs**
```
Summary:
  Total:        2.8745 secs
  Slowest:      0.0324 secs
  Fastest:      0.0003 secs
  Average:      0.0047 secs
  Requests/sec: 10436.6996
  Total data:   990000 bytes
  Size/request: 33 bytes
```

**Go 1 CPU**

```
Summary:
  Total:        3.8269 secs
  Slowest:      0.0508 secs
  Fastest:      0.0003 secs
  Average:      0.0063 secs
  Requests/sec: 7839.1591
  Total data:   990000 bytes
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

**Go 2 CPUs**

```
Summary:
  Total:        3.1866 secs
  Slowest:      0.0312 secs
  Fastest:      0.0003 secs
  Average:      0.0052 secs
  Requests/sec: 9414.3791
  Total data:   210000 bytes
  Size/request: 7 bytes
```

**Go 1 CPU**

```
Summary:
  Total:        4.8449 secs
  Slowest:      0.0422 secs
  Fastest:      0.0004 secs
  Average:      0.0080 secs
  Requests/sec: 6192.1232
  Total data:   210000 bytes
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