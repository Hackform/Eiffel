#!/usr/bin/env bash

docker run -it --rm --name cassserver -p 7000:7000 -p 7001:7001 -p 9042:9042 -v cassdata:/var/lib/cassandra cassandra:latest
