#!/bin/bash -ex
if [[ "$1" = "--data" ]]; then
	go-bindata data/
fi
go build
./go-rest --username docker --password docker --database docker --schema docker
