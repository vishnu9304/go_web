package main

import (
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	infoLog := log.New(os.Stdout, "INFO ", log.LUTC|log.Ldate|log.Lshortfile|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR ", log.LUTC|log.Ldate|log.Lshortfile|log.Ltime)

	// dependency injection
	app := application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	httpSrv := http.Server{
		Addr:    ":9999",
		Handler: app.loadRoutes(),
	}

	log.Fatal(httpSrv.ListenAndServe())
}
