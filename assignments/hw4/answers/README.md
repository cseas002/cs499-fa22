# Homework 4 Report

## Part 1

## Part 2

configuration:
node0		Workload generator
node1		Target server (case of single machine), Docker Swarm leader (case of multiple machines)
node2-4	Docker Swarm workers

command used
./wrk -t2 -c100 -d30s -R2000 -L -s ./scripts/hotel-reservation/mixed-workload_type_1.lua http://{node}:{port for mongodb-rate given by stack services}

Tutorial for using swarm with docker-compose (deploy stack)
https://docs.docker.com/engine/swarm/stack-deploy/

### MongoDB

#### Single node

#### Multi node

compare single to multi node

### Memcached
compare single to multi node

### Compare mongo to memcached