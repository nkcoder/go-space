package main

import (
	"net/http"
)

// This is an effective and idiomatic way to make dependencies available to our handlers
// without resorting to global variables or closures — any dependency that the
// healthcheckHandler needs can simply be included as a field in the application struct when
// we initialize it in main().
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writeJSON(w, 200, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
