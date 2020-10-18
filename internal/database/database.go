package database

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func syncToDisk() {

}

type databaseResponse struct {
}

func Select(res http.ResponseWriter, req *http.Request) {
	contextLogger := log.WithFields(log.Fields{"function": "database.Select"})
	contextLogger.Debug("handling /select request")
	res.WriteHeader(200)
	res.Write([]byte("select"))
}

func Insert(res http.ResponseWriter, req *http.Request) {
	contextLogger := log.WithFields(log.Fields{"function": "database.Insert"})
	contextLogger.Debug("handling /insert request")
	res.WriteHeader(200)
	res.Write([]byte("insert"))
}

func Update(res http.ResponseWriter, req *http.Request) {
	contextLogger := log.WithFields(log.Fields{"function": "database.Update"})
	contextLogger.Debug("handling /update request")
	res.WriteHeader(200)
	res.Write([]byte("update"))
}

func Delete(res http.ResponseWriter, req *http.Request) {
	contextLogger := log.WithFields(log.Fields{"function": "database.Delete"})
	contextLogger.Debug("handling /delete request: ")
	res.WriteHeader(200)
	res.Write([]byte("delete"))
}
