#!/bin/bash

#scion-bat
while true; do
    docker exec scion31 scion-bat http://16-ffaa:1:14,127.0.0.1:32765/hello &&
    echo
    sleep 0.5
done;






#docker exec scion31 scion-bat http://17-ffaa:1:21,127.0.0.1:32765/hello
