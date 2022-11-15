# Part 2: Evaluate Hotel Map Microservices

## Single node

Command used from node1
```
./wrk -t2 -c100 -d30s -R2000 -L -s ./scripts/hotel-reservation/mixed-workload_type_1.lua http://node0:8080 
```

### Monolithic Implementation
```
Running 30s test @ http://node0:8080
  2 threads and 100 connections
  Thread calibration: mean lat.: 0.489ms, rate sampling interval: 10ms
  Thread calibration: mean lat.: 0.476ms, rate sampling interval: 10ms
  Thread Stats   Avg      Stdev     99%   +/- Stdev
    Latency   482.58us  192.65us   1.04ms   74.72%
    Req/Sec     1.05k    86.32     1.22k    82.25%
#[Mean    =        0.910, StdDeviation   =        0.721]
#[Max     =        8.448, Total count    =        39896]
#[Buckets =           27, SubBuckets     =         2048]
----------------------------------------------------------
  58912 requests in 30.00s, 13.37MB read
  Non-2xx or 3xx responses: 23356
Requests/sec:   1963.55
Transfer/sec:    456.34KB
```

### Microservices Implementation
```
Running 30s test @ http://node0:8080
  2 threads and 100 connections
  Thread calibration: mean lat.: 1.769ms, rate sampling interval: 10ms
  Thread calibration: mean lat.: 1.747ms, rate sampling interval: 10ms
  Thread Stats   Avg      Stdev     99%   +/- Stdev
    Latency     1.77ms    1.48ms   4.42ms   68.62%
    Req/Sec     1.05k   154.98     1.44k    73.45%
#[Mean    =        1.774, StdDeviation   =        1.482]
#[Max     =       33.984, Total count    =        39897]
#[Buckets =           27, SubBuckets     =         2048]
----------------------------------------------------------
  59893 requests in 30.00s, 13.54MB read
  Non-2xx or 3xx responses: 24038
Requests/sec:   1996.52
Transfer/sec:    462.29KB

Search Nearby needed ~2ms and GetProfiles ~300us
```

As we can see, the monolithic implementation has much better latency and somewhat better throughput.
An expected result since the monolithic implementation is just a binary file whereas the microservices
implementation is bound by network speeds. It is worth mentioning that in, our case, where the containers
are all inside the same machine, network is probably not an issue, but the overhead of gRPC parsing
and other communivcation overheads that are causing the extra delays. 
