package config

import (
	"errors"
	"strings"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	_, err := LoadConfig(strings.NewReader(`
protocol: HTTP
host: 0.0.0.0
port: 23232
memoryLimit: ~
dataPersistence: true
certFile: self-signed.crt
keyFile: some-key.pem
logLevel: DEBUG
federate: localhost`))
	if err != nil {
		t.Fatal()
	}
}

func TestLoadInvalidConfig(t *testing.T) {
	_, err := LoadConfig(strings.NewReader(`invalid`))
	if err == nil {
		t.Fatal()
	}
}

type brokenReader struct{}

func (r brokenReader) Read(o []byte) (n int, err error) {
	return 0, errors.New("broke")
}

func TestLoadIOError(t *testing.T) {
	_, err := LoadConfig(brokenReader{})
	if err == nil {
		t.Fatal()
	}
}
