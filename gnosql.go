package main

import (
	"fmt"
	"io/ioutil"
	"net"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
)

type conf struct {
	Port        int64  `yaml:"Port"`
	MemoryLimit string `yaml:"MemoryLimit"`
}

func log(error error) {
	fmt.Print(error)
}

func executeDatabaseOperation() {

}

func handleRequest(conn net.Conn) {
	request, _ := conn.Read()
}

func server(config conf) {
	ln, err := net.Listen("tcp", ":"+string(config.Port))
	if err != nil {
		go log(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			go log(err)
		}
		go handleRequest(conn)
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

	go server(config)
}
