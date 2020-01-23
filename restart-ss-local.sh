#!/bin/bash

# kill ss-local first
self=$$
for i in `ps -ef | grep ss-local | awk '{print $2}' | sort -n`
do
    if [ $self != $i ]; then
        kill $i;
    fi
done

# start a new ss-local in background
nohup $1 -c $2 >> /dev/null &