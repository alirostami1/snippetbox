package main

import "net/http"

func (app *Application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// Static fileserving from ./ui/static
	fileserver := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	return mux
}
