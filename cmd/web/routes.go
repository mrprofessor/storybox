package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// File server that serves the static files in ui/static folder.
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// Register the static/ route where the static files can be served.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
