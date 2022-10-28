#!/bin/bash

# This script runs hotelapp for 1 to 4 hosts and gets results
for hosts in {1..4}
do
    bash configure_and_run_hotelapp.sh $hosts | grep -A 8 ^"Running 5s test @ http://node0" > ../Results/HotelApp/${hosts}results.txt
done

bash kill_hotelapp.sh # Kill the hotelapp in all possible hosts