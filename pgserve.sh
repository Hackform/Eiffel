#!/usr/bin/env bash

docker run -it --rm --name pgserver -p 5432:5432 -v pgdata:/var/lib/postgresql/data -e POSTGRES_PASSWORD=admin postgres:alpine
