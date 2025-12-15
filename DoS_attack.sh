#!/bin/bash
node1=scion23
sleep 2 &&
docker exec monitor scionctl capture start $node1 &&
docker exec scion23 tcpreplay -i eth1 -t --loop=100000 /etc/scion/ping_to_23_2.pcap &&
#sleep 5 &&
docker exec monitor scionctl capture stop $node1 &&

# save capture list to file
docker exec monitor scionctl capture list $node1 > captures/capture_list.txt &&
echo "Capture list saved to captures/capture_list.txt"

# Exporting the latest capture file
filename=$(awk '/\.pcap/ {print $4}' captures/capture_list.txt | tail -n1)
basename="${filename%.pcap}"
id=$(awk '/\.pcap/ {print $2}' captures/capture_list.txt | tail -n1)
baseid="${id%}"
docker exec monitor scionctl capture file $node1 $basename > captures/capture$baseid.pcap
echo "Capture file exported to captures/capture$baseid.pcap"

#copy file to container
#docker cp ping_to_21.pcap scion24:/etc/scion/

## works on node 23, but is not forwarded to node 21

##from 31 to 25 -> two paths cross 23, two don't
## 1. do scionping from 31 to 25
## 2. do scionping while tcp replay on 24 to 23
## 3. blacklist 22 on 31
## 4. do scionping with DoS from 31 to 25 again to see effect