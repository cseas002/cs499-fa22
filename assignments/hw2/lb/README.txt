hotelApp.sh does the whole job; It configures ansible with 1 to 4 webservers and gets result for all circumstances (1 node, 2 nodes, etc.)
The same goes with staticSite.sh

configure_and_run_hotelapp.sh configures hotelapp and ansible to run and get results for the hosts given in the arguments 
e.g. configure_and_run_hotelapp.sh 2 = It gets results for hotelapp with 2 webservers

run_hotelapp.sh runs the hotelapp to the 4 nodes. The mono file only

ansible_static_init.sh and hotelapp_init initialize ansible for the static site and hotelapp respectively

ansible_hosts_conf.sh configures the hosts-lb file with the hosts given in the argument 
e.g. ansible_hosts_config.sh 2 = with 2 webservers

