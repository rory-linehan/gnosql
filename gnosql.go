package main

import (
	"fmt"
	"io/ioutil"
	"net"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
)

type conf struct {
	NativePort  int64  `yaml:"NativePort"`
	APIPort     int64  `yaml:"APIPort"`
	MemoryLimit string `yaml:"MemoryLimit"`
}

func log(error error) {
	fmt.Print(error)
}

func executeDatabaseOperation() {

}

func handleAPIRequest(conn net.Conn) {

}

func handleNativeRequest(conn net.Conn) {

}

func server(mode string, config conf) {
	if mode == "native" {
		ln, err := net.Listen("tcp", ":"+string(config.NativePort))
		if err != nil {
			go log(err)
		}
		for {
			conn, err := ln.Accept()
			if err != nil {
				go log(err)
			}
			go handleNativeRequest(conn)
		}
	}
	if mode == "api" {
		ln, err := net.Listen("tcp", ":"+string(config.APIPort))
		if err != nil {
			go log(err)
		}
		for {
			conn, err := ln.Accept()
			if err != nil {
				go log(err)
			}
			go handleAPIRequest(conn)
		}
	}
}

func getConfig() conf {
	var c conf
	yamlFile, err := ioutil.ReadFile("etc/config.yaml")
	if err != nil {
		go log(err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		go log(err)
	}
	return c
}

func main() {
	config := getConfig()
	spew.Dump(config)

	go server("native", config)
	go server("api", config)
}
