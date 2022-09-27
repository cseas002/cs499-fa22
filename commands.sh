#!/bin/bash
cd 
sudo apt-get -y update 
ssh node1 sudo apt-get -y update
sudo cp ./assignments/hw1/metrics/DRAM/Linux/mlc /bin/
sudo apt-get -y install ansible
