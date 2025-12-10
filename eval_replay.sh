#!/bin/bash
node1=scion31
node2=scion23
docker exec monitor scionctl capture start $node1 &&
docker exec monitor scionctl scionping start $node2 $node1 &&
for i in {1..100}; do
    echo "Replaying capture from $node2 iteration $i"
  docker exec $node2 tcpreplay -i eth0 /etc/scion/capture10.pcap
done &&
docker exec monitor scionctl scionping stop $node2 &&
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