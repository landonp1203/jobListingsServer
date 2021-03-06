package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

const IDPath = "/lpatmore"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	s := r.PathPrefix(IDPath).Subrouter()
	s.HandleFunc("/all", GetAllDBRowsHandler).Methods("GET")
	s.HandleFunc("/status", GetDBStatusHandler).Methods("GET")

	_ = http.ListenAndServe(":8080", r)
}
