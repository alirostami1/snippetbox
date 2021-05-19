package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type Application struct {
	errorlog *log.Logger
	infolog  *log.Logger
}

func main() {
	// Command line flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initiate a new Application instance
	app := &Application{
		errorlog: errorLog,
		infolog:  infoLog,
	}

	// server mux routes

	// define the server
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// start the server
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
