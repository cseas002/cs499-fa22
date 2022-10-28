#!/bin/bash
for node in {1..4}
do
    ssh -A node$node git clone git@github.com:cseas002/cs499-fa22.git # Copy my repository in order to run the hotelapp
    ssh node$node "cd cs499-fa22; bash commands.sh" # Download important tools
    ssh -f node$node "cd /users/cseas002/cs499-fa22/labs/05-hotelapp/hotelapp; ./mono" # Run the hotelapp commands in the background (-f) 
done