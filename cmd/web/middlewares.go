package main

import (
	"net/http"
)

/*
	middleware syntax

	func someMiddleware(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// code block to be executed during forward flow
			next.ServerHttp(w, r)
			// code block to be executed during reverse flow
		})
	}

	Flow control:
	httpServer <-> Middleware <-> Mux <-> ApplicationHandler

*/

func (a application) logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.infoLog.Println(r.Method, r.Host, r.Proto, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
