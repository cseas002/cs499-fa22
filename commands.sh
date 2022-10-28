#!/bin/bash
cd 
sudo apt-get -y update 
ssh node1 sudo apt-get -y update
sudo cp ./assignments/hw1/metrics/DRAM/Linux/mlc /bin/
sudo apt-get -y install ansible
git clone https://github.com/wg/wrk.git
cd wrk
make -j
sudo cp wrk /bin/
git clone https://github.com/giltene/wrk2.git
cd wrk2
make -j
sudo cp wrk /bin/wrk2