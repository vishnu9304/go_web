package main

import (
	"net/http"
)

func (a application) ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func (a application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
