package gnosql

import (
	"github.com/davecgh/go-spew/spew"
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
	c := config.LoadConfig()
	spew.Dump(c)

	contextLogger.Info("setting log output level from config.LogLevel")
	switch c.LogLevel {
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
