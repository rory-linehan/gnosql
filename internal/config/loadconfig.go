package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Conf struct {
	Protocol        string `yaml:"protocol"`
	Address         string `yaml:"host"`
	Port            int64  `yaml:"port"`
	MemoryLimit     string `yaml:"memoryLimit"`
	DataPersistence bool   `yaml:"dataPersistence"`
	CertFile        string `yaml:"certFile"`
	KeyFile         string `yaml:"keyFile"`
	LogLevel        string `yaml:"logLevel"`
}

func LoadConfig() Conf {
	contextLogger := log.WithFields(log.Fields{"function": "getConfig"})
	contextLogger.Info("loading config")
	var c Conf
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		contextLogger.Fatal("failed to load config.yaml: " + err.Error())
	}
	contextLogger.Info("unmarshaling yaml from config")
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		contextLogger.Fatal("failed to unmarshal yaml: " + err.Error())
	}
	return c
}
