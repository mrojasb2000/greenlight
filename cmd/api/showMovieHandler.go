package main

import (
	"fmt"
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the application status, operating environment and version.

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	version := "0.0.1"
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}
