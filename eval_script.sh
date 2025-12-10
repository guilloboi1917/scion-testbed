#!/bin/bash
docker exec monitor scionctl capture start scion31 &&
docker exec monitor scionctl scionping start scion31 scion31 --count 10 &&
sleep 2 &&
docker exec monitor scionctl scionping stop scion31 &&
#docker exec scion31 scion-bat http://17-ffaa:1:23,[127.0.0.1]:32765/hello &&
docker exec monitor scionctl capture stop scion31 &&
docker exec monitor scionctl capture list scion31 > captures/capture_list.txt &&
echo "Capture list saved to captures/capture_list.txt"
filename=$(awk '/\.pcap/ {print $4}' captures/capture_list.txt | tail -n1)
basename="${filename%.pcap}"
id=$(awk '/\.pcap/ {print $2}' captures/capture_list.txt | tail -n1)
baseid="${id%}"
docker exec monitor scionctl capture file scion31 $basename > captures/capture$baseid.pcap

#echo "use 'docker exec monitor scionctl capture file scion31 capture_1765376986 > capture10.pcap' to export file"
