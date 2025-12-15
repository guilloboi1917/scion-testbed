#!/bin/bash
AS_1=scion31
AS_2=scion11

#scionping
while true; do
    docker exec $AS_1 scion traceroute 16-ffaa:1:11,10.100.0.11 &&
    sleep 3
done;