package main

import (
	"net/http"
)

func (a application) loadRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", a.ping)
	mux.HandleFunc("/home", a.home)
	return a.logRequests(mux)
}
