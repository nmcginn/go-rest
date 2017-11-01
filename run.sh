#!/bin/bash -ex
go build
./go-rest --username docker --password docker --database docker --schema docker
