#!/bin/bash -ex
docker build -t test-postgres .
docker run -it --rm -e POSTGRES_PASSWORD="passwordlmao" test-postgres
