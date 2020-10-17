package server

import (
	"github.com/jyro-io/gnosql/internal/config"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
)

func handleRequest(conn net.Conn) {

}

func Serve(c config.Conf) {
	contextLogger := log.WithFields(log.Fields{"function": "Serve"})
	switch c.Protocol {
	case "HTTPS":
		contextLogger.Info("starting HTTPS server")
		http.ListenAndServeTLS(c.Address+":"+string(c.Port), c.CertFile, c.KeyFile, nil)
	default:
		contextLogger.Info("starting default HTTP server")
		http.ListenAndServe(c.Address+":"+string(c.Port), nil)
	}
}
