# Benchmark Results

[boom](https://github.com/tarekziade/boom) was used as benchmarking tool.

See the benchmark results below.


## Decode JSON

The endpoint decode a JSON string and response with plain string

command : `boom -n 10000 -c 20 -m POST -D '{"name":"John", "username": "Doe"}'  http://localhost:5000/users`
**Go (2 CPUs)**

```
-------- Results --------
Successful calls                10000
Total time                      35.3489 s  
Average                         0.0626 s  
Fastest                         0.0280 s  
Slowest                         0.1755 s  
Amplitude                       0.1475 s  
Standard deviation              0.013287
RPS                             282
BSI                             Pretty good
```

**Go (1 CPU)**

```
-------- Results --------
Successful calls                10000
Total time                      36.3005 s  
Average                         0.0644 s  
Fastest                         0.0182 s  
Slowest                         0.1644 s  
Amplitude                       0.1462 s  
Standard deviation              0.012590
RPS                             275
BSI                             Pretty good
```

**Nim**
```
-------- Results --------
Successful calls                10000
Total time                      36.5024 s  
Average                         0.0648 s  
Fastest                         0.0191 s  
Slowest                         0.3244 s  
Amplitude                       0.3053 s  
Standard deviation              0.022291
RPS                             273
BSI                             Pretty good
```
