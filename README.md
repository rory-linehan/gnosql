# gnosql

gnosql is a dead simple, schema-free, cloud-native database system written
in Go.

## Overview

gnosql is designed to be the simplest, leanest, and most versatile open source
schema-free database system out there.

It's Docker and cloud native, designed to work with ephemeral requirements
as well as it does with reliability. It operates well under any memory
constraint, adapting its approach to management according to what resources are
available.

### Clients

Currently the following clients are supported:

- Go
- REST API

### Concurrency

Concurrency is the name of the game when developing applications of the future.
gnosql will handle any and all connections and concurrent operations as long as
you have the hardware to support it.

### Security

gnosql does not do SSL natively. If you want SSL (and you should), use `nginx`
and configure it as a reverse proxy.

### Operating Systems

gnosql is tested on Linux, if you want to use other operating systems,
you're on your own.

## Installation

#### Docker:

``

#### Go client import:

`import "github.com/rory-linehan/gosql/clients/gnosql-client"`

## Features

-

## Roadmap

-
