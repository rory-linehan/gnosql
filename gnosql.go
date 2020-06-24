package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
)

type conf struct {
	Protocol    string `yaml:"Protocol"`
	Address     string `yaml:"Address"`
	Port        int64  `yaml:"Port"`
	MemoryLimit string `yaml:"MemoryLimit"`
	Persistence bool   `yaml:"Persistence"`
	CertFile    string `yaml:"CertFile"`
	KeyFile     string `yaml:"KeyFile"`
}

func executeDatabaseOperation() {

}

func handleRequest(conn net.Conn) {

}

func server(config conf) {
	if config.Protocol == "HTTPS" {
		http.ListenAndServeTLS(
			config.Address+":"+string(config.Port),
			config.CertFile,
			config.KeyFile,
			nil)
	} else if config.Protocol == "HTTP" {
		http.ListenAndServe(config.Address+":"+string(config.Port), nil)
	} else {

	}
}

func getConfig() conf {
	var c conf
	yamlFile, err := ioutil.ReadFile("etc/config.yaml")
	if err != nil {
		//go log(err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		//go log(err)
	}
	return c
}

func main() {
	config := getConfig()
	spew.Dump(config)
	go server(config)
}
