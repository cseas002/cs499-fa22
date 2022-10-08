#!/bin/bash

./index.sh -i hosts-$1-index deploy
./index.sh -i hosts-$1-index start
./frontend.sh -i hosts-$1-index config
./frontend.sh -i hosts-$1-index start
./frontend.sh -i hosts-$1-index test
