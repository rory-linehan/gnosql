#!/usr/bin/env bash

version=$(cat VERSION)

go test -v -coverprofile cover.out ./... && \
./coverage.sh && \
go build -o bin/gnosql cmd/gnosql/main.go && \
sudo chmod +x bin/gnosql && \
./test/integration.sh && \
./test/integration.py && \
docker build -t gnosql:${version} . && \
docker tag gnosql:${version} gnosql:latest
