# Benchmark Results

[boom](https://github.com/tarekziade/boom) was used as benchmarking tool.

See the benchmark results below.

## Plaintext

The server endpoint simply return `Hello World!` string.

boom command : `boom -n 10000 -c 20 http://localhost:5000` which means:

- GET request
- number of request : 10K
- number of concurrent request : 20

**Go (4 CPUs)**

```
-------- Results --------
Successful calls		10000
Total time        		41.5772 s  
Average           		0.0740 s  
Fastest           		0.0304 s  
Slowest           		0.1532 s  
Amplitude         		0.1228 s  
Standard deviation		0.008222
RPS               		240
BSI              		Pretty good
```

**Go (1 CPU)**

```
Successful calls		10000
Total time        		42.3954 s  
Average           		0.0754 s  
Fastest           		0.0224 s  
Slowest           		0.2500 s  
Amplitude         		0.2275 s  
Standard deviation		0.010894
RPS               		235
BSI              		Pretty good

```


**Nim**

```
-------- Results --------
Successful calls        10000
Total time              41.0454 s  
Average                 0.0731 s  
Fastest                 0.0221 s  
Slowest                 0.1231 s  
Amplitude               0.1011 s  
Standard deviation      0.007544
RPS                     243
BSI                     Pretty good

-------- Status codes --------
Code 200                10000 times.
```

## Returns JSON

The server endpoint returns JSON array

boom command : `boom -n 10000 -c 20 http://localhost:5000/users`

**Go (4 CPUs)**
```

-------- Results --------
Successful calls		10000
Total time        		41.6466 s  
Average           		0.0742 s  
Fastest           		0.0297 s  
Slowest           		0.2710 s  
Amplitude         		0.2413 s  
Standard deviation		0.010528
RPS               		240
BSI              		Pretty good

```

**Go (1 CPU)**

```
Successful calls		10000
Total time        		42.3756 s  
Average           		0.0754 s  
Fastest           		0.0485 s  
Slowest           		0.3153 s  
Amplitude         		0.2668 s  
Standard deviation		0.012350
RPS               		235
BSI              		Pretty good
```

**Nim**

```
Successful calls        10000
Total time              41.0507 s  
Average                 0.0732 s  
Fastest                 0.0267 s  
Slowest                 0.1213 s  
Amplitude               0.0946 s  
Standard deviation      0.007702
RPS                     243
BSI                     Pretty good

-------- Status codes --------
Code 200                10000 times.
```

## Decode JSON

The endpoint decode a JSON string and response with plain string

command : `boom -n 10000 -c 20 -m POST -D '{"name":"John", "username": "Doe"}'  http://localhost:5000/users`
**Go (4 CPUs)**

```
Successful calls		10000
Total time        		41.8808 s  
Average           		0.0745 s  
Fastest           		0.0288 s  
Slowest           		0.1240 s  
Amplitude         		0.0951 s  
Standard deviation		0.007062
RPS               		238
BSI              		Pretty good
```

**Go (1 CPU)**

```
Successful calls		10000
Total time        		41.4922 s  
Average           		0.0738 s  
Fastest           		0.0253 s  
Slowest           		0.1177 s  
Amplitude         		0.0924 s  
Standard deviation		0.007345
RPS               		241
BSI              		Pretty good
```

**Nim**
```
-------- Results --------
Successful calls        10000
Total time              41.0145 s  
Average                 0.0730 s  
Fastest                 0.0253 s  
Slowest                 0.1241 s  
Amplitude               0.0988 s  
Standard deviation      0.007784
RPS                     243
BSI                     Pretty good

-------- Status codes --------
Code 200                10000 times.
```
