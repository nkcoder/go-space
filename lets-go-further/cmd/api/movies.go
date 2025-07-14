package main

import (
	"database/sql"
	"errors"
	"greenlight.danielguo.com/internal/data"
	"greenlight.danielguo.com/internal/validator"
	"net/http"
)

// For decoding JSON from HTTP request body, json.Decoder is generally the best choice. It's more efficient than
// json.Unmarshal, requires less code and offers some helpful settings to tweak its behavior.
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	// Go's http.Server will close r.Body automatically
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Copy the values from input to Movie
	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}
	// We might need to call multiple validation helpers
	v := validator.New()
	if movie.ValidateMovie(v); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Movies.Insert(movie)
	if err != nil {
		app.logger.Println("insert movie error:", err)
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"data": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie, err := app.models.Movies.Get(id)
	if err != nil {
		app.logger.Println("get movie from database error:", err)
		if errors.Is(err, sql.ErrNoRows) {
			http.NotFound(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"data": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
