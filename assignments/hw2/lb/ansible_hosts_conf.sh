#!/bin/bash

# Configures hosts with the number of arguments 
# e.g. bash ansible_hosts_conf.sh 2 - it configures ansible with 2 hosts as webservers

if [ $# -eq 0 ]
  then
    echo "No arguments supplied"
    exit 1
fi
temp=1
printf "[loadbalancer]\nnode0\n\n[webservers]\n" > hosts-lb
while true
do
	echo node$temp >> hosts-lb
	if [ $temp -ge $1 ]
	then
		break
	fi
	temp=$((temp+1))
done

echo "Configured ansible with 1 load balancer (node0) and $1 webservers"
