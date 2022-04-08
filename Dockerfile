FROM golang:1.17-alpine

RUN addgroup -S gnosql && adduser -S gnosql -G gnosql
RUN apk add --no-cache git
WORKDIR /home/gnosql
USER gnosql
ADD . .
RUN go build -o gnosql cmd/gnosql/main.go
RUN chmod +x gnosql
ENTRYPOINT ./gnosql --log INFO