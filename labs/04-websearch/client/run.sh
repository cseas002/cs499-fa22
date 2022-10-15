#!/bin/bash

# Script for running multiple threads and getting the average output
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
./client node1 8080 /local/websearch/ISPASS_PAPER_QUERIES_100K 1000 $thread onlyHits.jsp 1 1 /tmp/out 1 2> /dev/null  > temp.txt ; for node in {2..5}; do ssh node$node iostat | grep -A 1 ^avg >> node${node}_cpu.txt; done
cat temp.txt | grep ^"ops" > ops_temp.txt
cat temp.txt | grep ^"average" > avg_temp.txt
# Run the client and save average the results in a temp file and then save the ops/sec and average response latency to a temporary results file
awk '{print $2}' avg_temp.txt >> avg_results.txt # Remove the "average" text
awk '{print $2}' ops_temp.txt >> ops_results.txt # Remove the "ops/sec" text
done
rm temp.txt
rm ops_temp.txt
rm avg_temp.txt
