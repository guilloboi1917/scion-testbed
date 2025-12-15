#!/bin/bash
node2=scion23
docker exec monitor scionctl capture start $node2 &&
for i in {1..500}; do
docker exec scion31 scion-bat http://17-ffaa:1:23,127.0.0.1:32765/hello
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