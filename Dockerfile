FROM golang:1.13-alpine

RUN addgroup -S gnosql && adduser -S gnosql -G gnosql
RUN apk add --no-cache git
RUN go get gopkg.in/yaml.v2 github.com/davecgh/go-spew/spew
WORKDIR /home/gnosql
USER gnosql
ADD gnosql.go .
RUN go build -o gnosql
RUN chmod +x gnosql
ADD config.yaml .
ENTRYPOINT ./gnosql