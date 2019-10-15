#!/usr/bin/env bash

mkdir bin
go get -u github.com/davecgh/go-spew/spew
go get -u gopkg.in/yaml.v2
go build -o bin/gnosql

#docker build -t gnosql:latest .
