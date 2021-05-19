package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *Application) serverError(w http.ResponseWriter, err error) {
	// make add the goroutine trace to error log message
	trace := fmt.Sprintf("%s\n%s", err, debug.Stack())
	app.errorlog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
