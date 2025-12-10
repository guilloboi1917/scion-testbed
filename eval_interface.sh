#!/bin/bash
docker exec monitor scionctl capture start scion31 &&
sleep 1 &&
docker exec monitor scionctl capture start scion31 &&
sleep 1 &&
docker exec monitor scionctl capture stop scion31

# save capture list to file
docker exec monitor scionctl capture list scion31 > captures/capture_list.txt &&
echo "Capture list saved to captures/capture_list.txt"

# Exporting the latest capture file
filename=$(awk '/\.pcap/ {print $4}' captures/capture_list.txt | tail -n1)
basename="${filename%.pcap}"
id=$(awk '/\.pcap/ {print $2}' captures/capture_list.txt | tail -n1)
baseid="${id%}"
docker exec monitor scionctl capture file scion31 $basename > captures/capture$baseid.pcap
echo "Capture file exported to captures/capture$baseid.pcap"