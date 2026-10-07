package main

import (
    "net/http"      // HTTP error functions and status codes
    "runtime/debug" // Access runtime diagnostics (like stack traces)
)

// serverError logs the error details, logs the call stack trace,
// and sends a generic 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
    var (
        method = r.Method               // GET, POST, etc.
        uri    = r.URL.RequestURI()     // The path requested by the client
        trace  = string(debug.Stack())  // The exact line numbers where the code broke
    )

    // Log the error via structured slog with helpful metadata attributes.
    app.logger.Error(err.Error(), "method", method, "uri", uri, "trace", trace)

    // Send a 500 error to the client with clean text ("Internal Server Error").
    // We do NOT send raw errors to users to avoid leaking internal system details.
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError sends a specific client error code (e.g. 400 Bad Request, 405 Method Not Allowed).
func (app *application) clientError(w http.ResponseWriter, status int) {
    http.Error(w, http.StatusText(status), status)
}