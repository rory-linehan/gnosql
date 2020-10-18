#!/usr/bin/env bash

version=$(cat VERSION)

go test -v -coverprofile cover.out ./... && \
./test/integration.sh && \
go build -o gnosql cmd/gnosql/main.go && \
sudo chmod +x gnosql && \
docker build -t gnosql:${version} . && \
docker tag gnosql:${version} gnosql:latest
