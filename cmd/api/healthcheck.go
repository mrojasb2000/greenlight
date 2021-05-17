package main

import (
	"net/http"
)

// Declare a handler which writes a plain-text response with information about
// the application status, operating environment and version.

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	version := "0.0.1"

	// we've constructed this means the environment and version data will now be nested
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	// Declare an envelope map containing the data for the response. Notice that the way

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
