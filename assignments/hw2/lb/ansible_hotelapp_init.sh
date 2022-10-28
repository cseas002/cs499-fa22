#!/bin/bash

ansible-playbook -i hosts-lb nginx-lb.yml
bash run_hotelapp.sh