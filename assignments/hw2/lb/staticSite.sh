#!/bin/bash

# This script runs staticSite for 1 to 4 hosts and gets results
for hosts in {1..4}
do
    bash ansible_hosts_conf.sh $hosts
    bash ansible_static_init.sh
    ansible-playbook -i hosts-lb nginx-lb.yml

    # Take wrk measurements
    ssh node5 "wrk -t500 -c500 -d5s --timeout 2s http://node0 2> /dev/null" > ../Results/Static/${hosts}results.txt
done