#!/usr/bin/env bash

go test -v -coverprofile cover.out ./...
go tool cover -html=cover.out -o coverage-unit.html