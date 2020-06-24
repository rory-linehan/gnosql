package main

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net"
	"net/http"
	"os"

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
	LogLevel    string `yaml:"LogLevel"`
}

func executeDatabaseOperation() {

}

func handleRequest(conn net.Conn) {

}

func server(config conf) {
	contextLogger := log.WithFields(log.Fields{"function": "server"})
	switch config.Protocol {
	case "HTTPS":
		contextLogger.Info("starting HTTPS server")
		http.ListenAndServeTLS(config.Address+":"+string(config.Port), config.CertFile, config.KeyFile, nil)
	default:
		contextLogger.Info("starting default HTTP server")
		http.ListenAndServe(config.Address+":"+string(config.Port), nil)
	}
}

func getConfig() conf {
	contextLogger := log.WithFields(log.Fields{"function": "getConfig"})
	contextLogger.Info("loading config.yaml")
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

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.ErrorLevel)
}

func main() {
	contextLogger := log.WithFields(log.Fields{"function": "main"})
	config := getConfig()
	spew.Dump(config)

	contextLogger.Info("setting log output level from config.LogLevel")
	switch config.LogLevel {
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	case "PANIC":
		log.SetLevel(log.PanicLevel)
	default:
		log.SetLevel(log.ErrorLevel)
	}

	server(config)
}
