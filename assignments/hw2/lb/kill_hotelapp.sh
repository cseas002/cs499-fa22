#!/bin/bash
for node in {1..5}
do
    ssh node$node pkill mono # kill the process
done
