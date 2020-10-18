# gnosql

gnosql is a schema-free, fully 
[HTTP](https://tools.ietf.org/html/rfc2616) 
compliant distributed database system written in Go.

### Overview

gnosql is designed to be the simplest, most versatile open source
schema-free database system in existence.

It is HTTP compliant, more specifically the interface for gnosql is simply
an HTTP compliant REST API. 
No more clunky client interfaces wrapping some obscure C library,
or extra packages that will break your software.
If you have code that handles HTTP requests,
then with gnosql you also have a database.

It's cloud native, designed to work with ephemeral requirements
as well as it does with reliability constraints.

Initial development is oriented toward ease-of-use and stability,
once things are stable the priority switches to efficiency.

### Usage

All requests must be POSTed JSON objects. 
You should expect JSON in response, 
along with an HTTP status code letting you know how to proceed.

### API

#### `/select`

#### `/insert`

#### `/update`

#### `/delete`