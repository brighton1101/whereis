// Package server contains a simple webserver that mirrors the cli's
// functionality.
package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/brighton1101/whereis/core"
)

const (
	WhereisPath = "/"
	TargetUriQueryParam = "uri"
	Port = ":5000"
)

func HandleWhereis(w http.ResponseWriter, r *http.Request) {
	inuri := r.URL.Query().Get(TargetUriQueryParam)
	if inuri == "" {
		w.WriteHeader(404)
		return
	}
	redirres, redirerr := core.GetNoRedirect(inuri)
	if redirerr != nil {
		w.WriteHeader(404)
		return
	}
	resstr, jserr := json.Marshal(redirres)
	if jserr != nil {
		w.WriteHeader(404)
		return
	}
	w.Write([]byte(resstr))
}


func Run() {
	http.HandleFunc(WhereisPath, HandleWhereis)
	log.Fatal(http.ListenAndServe(Port, nil))
}