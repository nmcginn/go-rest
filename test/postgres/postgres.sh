#!/bin/bash -ex
docker build -t test-postgres .
docker run -it --rm -p 5432:5432 -e POSTGRES_PASSWORD="passwordlmao" test-postgres
