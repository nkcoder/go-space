package main

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

// This method doesn't use any dependencies from our application struct so it could just be a regular function, rather
// than a method on application.
// But in general, I suggest setting up all your application-specific handlers and helpers so that they are methods on
// application. It helps maintain consistency in your code structure, and also future-proofs your code for when those
// handlers and helpers change later and they do need access to a dependency.
func (app *application) readIDParam(r *http.Request) (int64, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	// MarshalIndent has slight performance overhead than Marshal
	js, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(js)

	return nil
}

type envelope map[string]any
