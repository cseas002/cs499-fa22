#!/bin/bash

if [ $# -eq 0 ]
  then
    echo "No arguments supplied. Type the number of hosts"
    exit 1
fi

bash ansible_hosts_conf.sh $1
bash ansible_hotelapp_init.sh

# Take wrk measurements
ssh node5 "cd /users/cseas002/cs499-fa22/labs/05-hotelapp/hotelapp/wrk2; wrk2 -t100 -c100 -d5s -R99999999 -s ./scripts/hotel-reservation/mixed-workload_type_1.lua http://node0 2> /dev/null" 