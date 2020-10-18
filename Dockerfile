FROM golang:1.13-alpine

RUN addgroup -S gnosql && adduser -S gnosql -G gnosql
RUN apk add --no-cache git
WORKDIR /home/gnosql
USER gnosql
ADD . .
RUN go build cmd/gnosql/main.go -o gnosql
RUN chmod +x gnosql
ENTRYPOINT ./gnosql --log INFO