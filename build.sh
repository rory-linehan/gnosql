#!/usr/bin/env bash

version=$(cat VERSION)

go test -v -coverprofile cover.out ./... && \
./coverage.sh && \
go build -o gnosql.bin cmd/gnosql/main.go && \
sudo chmod +x gnosql.bin && \
./test/integration.sh && ./test/integration.py && \
rm gnosql.bin
docker build -t gnosql:${version} . && \
docker tag gnosql:${version} gnosql:latest
