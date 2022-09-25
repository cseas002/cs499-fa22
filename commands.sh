#!/bin/bash
cd 
git clone git@github.com:ucy-coast/cs499-fa22.git
sudo apt-get -y update 
ssh node1 sudo apt-get -y update
sudo cp ./assignments/hw1/metrics/DRAM/Linux/mlc /bin/
