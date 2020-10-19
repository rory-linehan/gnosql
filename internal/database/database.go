package database

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func syncToDisk() {

}

type databaseResponse struct {
	Code int
	Body []byte
}

func Select(res http.ResponseWriter, req *http.Request) {
	contextLogger := log.WithFields(log.Fields{"function": "database.Select"})
	contextLogger.Debug("handling /select request")
	resp := databaseResponse{
		Code: 200,
		Body: []byte("select"),
	}
	res.WriteHeader(resp.Code)
	res.Write(resp.Body)
}

func Insert(res http.ResponseWriter, req *http.Request) {
	contextLogger := log.WithFields(log.Fields{"function": "database.Insert"})
	contextLogger.Debug("handling /insert request")
	resp := databaseResponse{
		Code: 200,
		Body: []byte("insert"),
	}
	res.WriteHeader(resp.Code)
	res.Write(resp.Body)
}

func Update(res http.ResponseWriter, req *http.Request) {
	contextLogger := log.WithFields(log.Fields{"function": "database.Update"})
	contextLogger.Debug("handling /update request")
	resp := databaseResponse{
		Code: 200,
		Body: []byte("update"),
	}
	res.WriteHeader(resp.Code)
	res.Write(resp.Body)
}

func Delete(res http.ResponseWriter, req *http.Request) {
	contextLogger := log.WithFields(log.Fields{"function": "database.Delete"})
	contextLogger.Debug("handling /delete request: ")
	resp := databaseResponse{
		Code: 200,
		Body: []byte("delete"),
	}
	res.WriteHeader(resp.Code)
	res.Write(resp.Body)
}
