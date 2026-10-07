package main

import (
	"fmt"           // Used for string formatting (fmt.Fprintf)
	"html/template" // Safe HTML template engine that prevents XSS attacks
	"net/http"      // HTTP constants and types
	"strconv"       // String conversions (e.g. string to int)
)

// Home page handler
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Add a custom response header indicating the server software.
	w.Header().Add("Server", "Go")

	// List the template files needed to render the home page.
	// The base layout MUST be first so Go knows what to parse as the foundation[cite: 4].
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	// Parse the template files from disk into memory.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		// If file reading or parsing fails, call our centralized serverError helper.
		app.serverError(w, r, err)
		return
	}

	// Execute the "base" template and stream the rendered HTML into w (the browser).
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// View a specific snippet: /snippet/view/123
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// r.PathValue("id") extracts the {id} wildcard value defined in routes.go.
	// strconv.Atoi converts that string into an integer.
	id, err := strconv.Atoi(r.PathValue("id"))

	// Validation: ID must be a valid integer and greater than 0.
	if err != nil || id < 1 {
		// http.NotFound sends a standard 404 response back to the client.
		http.NotFound(w, r)
		return
	}

	// Write a confirmation response back to the browser.
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Show the HTML form for creating a snippet
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// w.Write takes a byte slice ([]byte) and writes it as the HTTP response body.
	w.Write([]byte("Display a form for creating a new snippet..."))
}

// Handle the submission of the snippet creation form (POST)
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Set HTTP status code 201 Created (default is 200 OK).
	// Note: w.WriteHeader must be called BEFORE w.Write.
	w.WriteHeader(http.StatusCreated)

	// Write response body
	w.Write([]byte("Save a new snippet..."))
}
