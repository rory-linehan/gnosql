package main

import (
	"flag"
	"github.com/jyro-io/gnosql/internal/config"
	"github.com/jyro-io/gnosql/internal/server"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.ErrorLevel)
}

func main() {
	contextLogger := log.WithFields(log.Fields{"function": "main"})

	var logLevel string
	var configFile string
	flag.StringVar(&logLevel, "log", "INFO", "log level [TRACE,DEBUG,INFO,WARN,ERROR,FATAL,PANIC]")
	flag.StringVar(&configFile, "config-file", "config.yaml", "config file")

	file, err := os.Open(configFile)
	if err != nil {
		contextLogger.Fatal("failed to read config file")
	}
	c, err := config.LoadConfig(file)
	if err != nil {
		contextLogger.Fatal("failed to load config from given file")
	}

	contextLogger.Info("setting log output level from config.LogLevel")
	switch logLevel {
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

	server.Serve(c)
}
