package server

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func NewHTTPServer(addr string) *http.Server {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		glob, _ := template.ParseGlob("./html/*")
		glob.ExecuteTemplate(w, "home.gohtml", nil)
	}).Methods("GET")

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
