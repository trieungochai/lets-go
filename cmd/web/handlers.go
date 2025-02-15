package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

// Change the signature of the home handler
// so it is defined as a method against *application.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	// Initialize a slice containing the paths to the two files.
	// It's important to note that
	// the file containing our base template must be the *first* file in the slice.
	files := []string{
		"./ui/html/base.tmpl",
		// Include the navigation partial in the template files.
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	// Use the template.ParseFiles() func to read the files and store the templates in a template set.
	// Notice that we use ... to pass the contents of the files slice as variadic arguments.
	// If there's an error, we log the detailed error message,
	// use the http.Error() function to send an Internal Server Error response to the user,
	// and then return from the handler so no subsequent code is executed.
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		// Because the home handler is now a method against the application struct
		// it can access its fields, including the structured logger.
		// We'll use this to create a log entry at Error level containing the error message,
		// also including the request method and URI as attributes to assist with debugging.
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		return
	}

	// Use the ExecuteTemplate() method to write the content of the "base" template as the response body.
	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		// And we also need to update the code here to use the structured logger too.
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Change the signature of the home handler
// so it is defined as a method against *application.
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Change the signature of the home handler
// so it is defined as a method against *application.
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

// Change the signature of the home handler
// so it is defined as a method against *application.
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet..."))
}
