package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Handle all not defined routes
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	// Use template.ParseFiles to read the template files into a template set.
	htmlFiles := []string{
		"./ui/html/home.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}
	ts, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Execute the template set to write the template content as the response
	// body. We can send dynamic data in the Execute method as the last
	// parameter.
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w , "Display a specific snippet with ID %d\n", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet...\n"))
}
