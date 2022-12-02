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

#### Throughput
![alt text](https://github.com/cseas002/cs499-fa22/blob/main/assignments/hw4/answers/results/throughput.png "Throughput")

