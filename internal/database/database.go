package database

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func syncToDisk() {

}

func Select(res http.ResponseWriter, req *http.Request) {
	contextLogger := log.WithFields(log.Fields{"function": "database.Select"})
	res.WriteHeader(200)
	res.Write([]byte("select!"))
}

func Insert(res http.ResponseWriter, req *http.Request) {
	contextLogger := log.WithFields(log.Fields{"function": "database.Insert"})
	res.WriteHeader(200)
	res.Write([]byte("insert!"))
}

func Update(res http.ResponseWriter, req *http.Request) {
	contextLogger := log.WithFields(log.Fields{"function": "database.Update"})
	res.WriteHeader(200)
	res.Write([]byte("update!"))
}

func Delete(res http.ResponseWriter, req *http.Request) {
	contextLogger := log.WithFields(log.Fields{"function": "database.Delete"})
	res.WriteHeader(200)
	res.Write([]byte("delete!"))
}
