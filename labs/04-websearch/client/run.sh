#!/bin/bash

# Script for running multiple threads and getting the average output
for thread in {1..128} # For 1 to <argument value> threads
do
echo "Running with $thread threads"
./client node1 8080 /local/websearch/ISPASS_PAPER_QUERIES_100K 1000 $thread onlyHits.jsp 1 1 /tmp/out 1 2> /dev/null | grep ^"average" > temp.txt 
# Run the client and save average the results in a temp file
awk '{print $2}' temp.txt >> results.txt # Remove the "average" text
done
rm temp.txt

