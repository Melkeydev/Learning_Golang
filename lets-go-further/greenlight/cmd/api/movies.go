package main

import (
  "fmt"
  "time"
  "net/http"
  "github.com/amokstakov/greenlight/internal/data"
  "github.com/amokstakov/greenlight/internal/validator"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
  var input struct {
    Title string `json:"title"`
    Year int32 `json:"year"`
    Runtime int32 `json:"runtime"`
    Genres []string `json:"genres"`
  }

  err := app.readJSON(w, r, &input)
  if err != nil {
    app.badRequestResponse(w, r, err)
    return
  }

  movie := &data.Movie{
    Title: input.Title,
    Year: input.Year,
    Runtime: input.Runtime,
    Genres: input.Genres,
  }

  v := validator.New()

  if data.ValidateMovie(v, movie); !v.Valid() {
    app.failedValidationResponse(w, r, v.Errors)
    return
  }

  fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
  id, err := app.readIDParam(r) 
  if err != nil {
    app.notFoundResponse(w, r)
    return
  }

  movie := data.Movie{
    ID: id,
    CreatedAt: time.Now(),
    Title: "Movie Title",
    Runtime: 102,
    Genres: []string{"drama", "romance", "war"},
    Version: 1,
  }

  err = app.writeJSON(w, http.StatusOK, envelope{"movie":movie}, nil)
  if err != nil {
    app.serverErrorResponse(w, r, err)
  }
}
