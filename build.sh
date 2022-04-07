#!/usr/bin/env bash

version=$(cat VERSION)

go test -v -coverprofile cover.out ./... && \
./coverage.sh && \
go build -o gnosql cmd/gnosql/main.go && \
sudo chmod +x gnosql && \
./test/integration.sh && ./test/integration.py && cd .. \
docker build -t gnosql:${version} . && \
docker tag gnosql:${version} gnosql:latest
