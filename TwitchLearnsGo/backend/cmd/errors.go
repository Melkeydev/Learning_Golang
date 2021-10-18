package main

import (
	"fmt"
	"net/http"
)
 
// Create a handler to log our errors
func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

// Helper for sending JSON-formatted error messages 
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil) 
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// Helper for handling unexpected server errors like 500 
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	// TODO: Change this lol
	message := "The server encountered an unexpected problem - sozz fam"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// Helper if client makes a request that does not exist
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "The requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// Helper when requested method is made incorrectle (get vs post)
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprint("the requested method %s is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusBadRequest, message)
}

// Helper for handling bad requests
func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

// handler for incorrect JSON validation
func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}























