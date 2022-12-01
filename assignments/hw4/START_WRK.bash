cd hotelapp
sudo apt-get install -y luarocks
sudo luarocks install luasocket
cd wrk2
make
./wrk -t2 -c100 -d30s -R2000 -L -s ./scripts/hotel-reservation/mixed-workload_type_1.lua http://node1:8080