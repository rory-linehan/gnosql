#!/usr/bin/env bash

go build -o gnosql cmd/gnosql/main.go && \
sudo chmod +x gnosql

#docker build -t gnosql:latest .
