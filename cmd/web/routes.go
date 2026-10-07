package main

import (
	"net/http" // Standard web library
)

// routes() is a method on *application. This allows it to access app.home,
// app.snippetView, etc., which need access to the database and logger.
func (app *application) routes() *http.ServeMux {
	// http.NewServeMux creates a new HTTP request multiplexer (router).
	mux := http.NewServeMux()

	// Set up a file server to serve static assets (CSS, images, JS) from ./ui/static/[cite: 3, 4]
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// When a request begins with "/static/", strip that prefix away so the file server
	// looks inside "./ui/static/" instead of "./ui/static/static/"[cite: 3, 4].
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Go 1.22+ routing syntax: "METHOD PATH"
	// "GET /{$}" means match EXACTLY the root path "/" (prevent matching /anything-else).
	mux.HandleFunc("GET /{$}", app.home)

	// Wildcard route: {id} captures the snippet ID from the URL path.
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)

	// Form display route
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)

	// Form submission route (handles POST requests to the same path)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	return mux // Return the configured router back to main.go
}
