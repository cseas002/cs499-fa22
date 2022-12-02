# Homework 4 Report

## Part 1

## Part 2

Configuration:
```
node0     Workload generator
node1     Target server (case of single machine), Docker Swarm leader (case of multiple machines)
node2-4   Docker Swarm workers
```

To generate workloads we used
```
./wrk -t2 -c100 -d30s -R2000 -L -s ./scripts/hotel-reservation/mixed-workload_type_1.lua http://{node}:{port for mongodb-rate given by stack services}
```

Tutorial for [using swarm with docker-compose (deploy stack)](https://docs.docker.com/engine/swarm/stack-deploy/)

### MongoDB

#### Single node
![alt text](https://github.com/cseas002/cs499-fa22/blob/main/assignments/hw4/answers/results/mongodb_single.png "MongoDB single node")

#### Multi node
![alt text](https://github.com/cseas002/cs499-fa22/blob/main/assignments/hw4/answers/results/mongodb_multi.png "MongoDB multi node") 

### Memcached

#### Single node
![alt text](https://github.com/cseas002/cs499-fa22/blob/main/assignments/hw4/answers/results/memcached_single.png "Memcached single node")

#### Multi node
![alt text](https://github.com/cseas002/cs499-fa22/blob/main/assignments/hw4/answers/results/memcached_multi.png "Memcached multi node") 

### Results analysis

#### Latency
![alt text](https://github.com/cseas002/cs499-fa22/blob/main/assignments/hw4/answers/results/latency.png "Latency")
Regarding latency, what we observe is that in both cases MongoDB and Memcached, the latency starts off small in the single node with a value of around 2 and 4 respectively, and there is a dramatic increase to around 300 and 400ms. This can be explained by the bottleneck that the network introduces and therefore is expected. What was not expected is the fact that the memcached latency is higher in both single and multinode implementations. To some degree, this can be explained by cache misses and the system trying to fill the cache with relevant values. However, with the workload we have, we would expect that by some point the cache would produce better average results than MongoDB. Maybe this difference in values is due to the implementations of each technology and not to the database and cache as concepts in a system. 

#### Throughput
![alt text](https://github.com/cseas002/cs499-fa22/blob/main/assignments/hw4/answers/results/throughput.png "Throughput")
Similarly here in throughput, both implementations MongoDB and Memcached start of well and have a drop in performance, although not as dramatic as in latency. Again, we can see that the Memcached implementation performs worse but in this case it is probably the expected result since the Memcached implementation does more operations per request and therefore takes longer for each request to get resolved.

#### Conclusion
Apparently, there is not a "good" or "bad" solution, either that being a database with a cache or not, or deplying to single or multiple nodes. Everything depends on the use case and we can only find the optimal implementation by testing

Thank you for reading :)
