package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mrojasb2000/greenlight/internal/data"
)

// Add a createMovieHandler for the "POST /v1/movies" endpoint. For now we simply
// return a plain-text plcaholder response

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

// Add a showMovieHandler for the "GET /v1/movies/:id" endpoint. For now, we retrive
// the interpolated "id" parameter from the current URL and include it in a placeholder
// response.

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Create a new instace of the Movie struct, containing the ID we extracted from
	// the URL and some dummt data. Alto notice that we deliberately haven't set a
	// value for the Year field.
	movie := data.Movie{
		ID:       id,
		CreateAt: time.Now(),
		Title:    "Casablanca",
		Runtime:  102,
		Genres:   []string{"drama", "romance", "war"},
		Version:  1,
	}

	// Encode the struct ro JSON and send it as the HTTP response
	err = app.writeJSON(w, http.StatusOK, movie, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}