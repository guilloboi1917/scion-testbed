#!/bin/bash
docker exec monitor scionctl capture start scion23 &&
docker exec monitor scionctl scionping start scion31 scion31 --count 10 &&
sleep 5 &&
docker exec monitor scionctl scionping stop scion31 &&
#docker exec scion31 scion-bat http://17-ffaa:1:23,[127.0.0.1]:32765/hello &&
docker exec monitor scionctl capture stop scion31 &&
docker exec monitor scionctl capture list scion31
#echo "use 'docker exec monitor scionctl capture file scion31 capture_1765376986 > capture10.pcap' to export file"
