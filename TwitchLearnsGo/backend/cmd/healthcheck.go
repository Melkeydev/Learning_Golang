package main

import (
	//"log"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// We need to creeate an envelope struct to hanlde our messaging

	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.Env,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil{
		app.serverErrorResponse(w, r, err)
	}
}
