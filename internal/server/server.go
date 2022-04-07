package server

import (
	"github.com/rory-linehan/gnosql/internal/config"
	"github.com/rory-linehan/gnosql/internal/database"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Serve(c config.Conf) {
	contextLogger := log.WithFields(log.Fields{"function": "server.Serve"})

	mux := http.NewServeMux()
	mux.HandleFunc("/select", database.Select)
	mux.HandleFunc("/insert", database.Insert)
	mux.HandleFunc("/update", database.Update)
	mux.HandleFunc("/delete", database.Delete)

	switch c.Protocol {
	case "HTTPS":
		contextLogger.Info("starting HTTPS server")
		contextLogger.Fatal(http.ListenAndServeTLS(c.Address+":"+c.Port, c.CertFile, c.KeyFile, mux))
	default:
		contextLogger.Info("starting default HTTP server")
		contextLogger.Fatal(http.ListenAndServe(c.Address+":"+c.Port, mux))
	}
}
