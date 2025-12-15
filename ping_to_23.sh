#!/bin/bash
node1=scion23
node2=scion24
docker exec monitor scionctl capture start $node2 &&
docker exec monitor scionctl scionping start $node2 $node1 &&
sleep 50 &&
docker exec monitor scionctl scionping stop $node2 &&
for i in {1..500}; do
docker exec $node2 scion-bat http://17-ffaa:1:23,127.0.0.1:32765/hello
done &&
docker exec monitor scionctl capture stop $node2 &&

# save capture list to file
docker exec monitor scionctl capture list $node2 > captures/capture_list.txt &&
echo "Capture list saved to captures/capture_list.txt"

# Exporting the latest capture file
filename=$(awk '/\.pcap/ {print $4}' captures/capture_list.txt | tail -n1)
basename="${filename%.pcap}"
id=$(awk '/\.pcap/ {print $2}' captures/capture_list.txt | tail -n1)
baseid="${id%}"
docker exec monitor scionctl capture file $node2 $basename > captures/capture$baseid.pcap
echo "Capture file exported to captures/capture$baseid.pcap"

#docker exec monitor scionctl capture file $node1 capture_1765385328 > capture10.pcap"

#copy file to node 24 and tcpreplay it
#docker cp ping_to_21.pcap scion24:/etc/scion/
#docker exec scion24 tcpreplay -i eth0 -t --loop=1000000 /etc/scion/ping_to_21.pcap