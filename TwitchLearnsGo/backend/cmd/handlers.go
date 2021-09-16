package main

import (
  "encoding/json"
  "net/http"
)

type AppStatus struct {
  Status string
}

func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {

  status := struct {
    Status string
  }{"Curling Teddy"}

  js, err := json.MarshalIndent(status, "", "\t")
  if err != nil {
    app.logger.Println(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(js)
}
