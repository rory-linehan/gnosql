# gnosql

gnosql is a schema-free, fully HTTP compliant database system written in Go.

### Overview

gnosql is designed to be the simplest, most versatile open source
schema-free database system in existence.

It is HTTP compliant, more specifically the interface for gnosql is simply
an HTTP compliant REST API. No more clunky client interfaces or extra
packages that will break your software.

It's cloud native, designed to work with ephemeral requirements
as well as it does with reliability. It operates well under any memory
constraint, adapting its approach to management according to what resources are
available.

### Concurrency

gnosql will handle any and all connections and concurrent operations as long as
you have the hardware to support it.
