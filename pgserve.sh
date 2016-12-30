#!/usr/bin/env bash

docker run -d --name pgserver -p 5432:5432 -v db:/var/lib/postgres/data -e POSTGRES_PASSWORD=admin postgres:alpine
