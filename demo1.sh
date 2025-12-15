#!/bin/bash
AS_1=scion31
AS_2=scion11

#scionping
while true; do
    docker exec $AS_1 sh -c 'rm /var/lib/scion-node-manager/scion-ping-results/*'
    docker exec monitor scionctl scionping start $AS_1 $AS_2  --count 1 &&
    sleep 2 &&
    docker exec $AS_1 sh -c 'cat /var/lib/scion-node-manager/scion-ping-results/*' &&
    sleep 4
done;
