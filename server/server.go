package server

import (
	"net/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer() *http.Server {
	r := mux.NewRouter()
	//r.HandleFunc("/").Methods("GET")
	//return &http.Server{
	//	Handler: r,
	//}
}
