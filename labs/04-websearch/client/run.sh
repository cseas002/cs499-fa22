#!/bin/bash

# Script for running multiple threads and getting the average output
export TERM=linux
for node in {2..5} 
do
	rm node${node}_cpu.txt
	ssh node$node sudo apt-get -y install sysstat
done
rm avg_results.txt
rm ops_results.txt
# Removing the old node cpu result files and installing sysstat to other nodes in order to execute iostat
for thread in {1..128} 
do
echo "Running with $thread threads"
./client node1 8080 /local/websearch/ISPASS_PAPER_QUERIES_100K 1000 $thread onlyHits.jsp 1 1 /tmp/out 1 2> /dev/null > temp.txt & 
# This will run in the background in order to read the CPU utilization at the same time
for node in {2..5}
do 
	# ssh node$node export TERM=xterm-256color
	# ssh -t node$node top -o %CPU -n 1 | grep %CPU -A 1 | tail -1 | awk '{print $9}' >> node${node}_cpu.txt 
	ssh node$node mpstat | grep CPU -A 1 | tail -1 | awk '{print $4}' >> node${node}_cpu.txt 
done
# For every node from 2 to 5 it will measure the CPU throughput and append it in a file named "node<node number>_cpu.txt"
cat temp.txt | grep ^"ops" > ops_temp.txt
cat temp.txt | grep ^"average" > avg_temp.txt
# Run the client and save average the results in a temp file and then save the ops/sec and average response latency to a temporary results file
awk '{print $2}' avg_temp.txt >> avg_results.txt # Remove the "average" text
awk '{print $2}' ops_temp.txt >> ops_results.txt # Remove the "ops/sec" text
done
rm temp.txt
rm ops_temp.txt
rm avg_temp.txt
