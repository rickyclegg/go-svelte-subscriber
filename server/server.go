package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewHTTPServer(addr string) *http.Server {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("foo"))
	}).Methods("GET")

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
