#!/usr/bin/env bash

version=$(cat VERSION)

./test/integration.sh && \
go build -o gnosql cmd/gnosql/main.go && \
sudo chmod +x gnosql && \
docker build -t gnosql:${version} . && \
docker tag gnosql:${version} gnosql:latest
