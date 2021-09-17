package main 

import (
  //"errors"
  "log"
  "strconv"
  "net/http"
  "encoding/json"
  "backend/models"
  "github.com/julienschmidt/httprouter"
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

func (app *application) getJob(w http.ResponseWriter, r *http.Request) {
  params := httprouter.ParamsFromContext(r.Context())

  id, err := strconv.Atoi(params.ByName("id"))
  if err != nil {
    app.logger.Print("invalid id parameter")
    return
  }

  app.logger.Print(id)

  user, err := app.models.DB.GetJob(id)
  if err != nil {
    app.logger.Println("Could not get job")
  }

  js, err := json.MarshalIndent(user, "", "\t")
  if err != nil {
    app.logger.Println(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(js)

}

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {
  params := httprouter.ParamsFromContext(r.Context())

  id, err := strconv.Atoi(params.ByName("id"))
  if err != nil {
    app.logger.Print("invalid id parameter")
    return
  }

  app.logger.Print(id)

  user, err := app.models.DB.GetUser(id)
  if err != nil {
    app.logger.Println("Could not get user")
  }

  js, err := json.MarshalIndent(user, "", "\t")
  if err != nil {
    app.logger.Println(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(js)
}

type UserPayload struct {
  Username string `json:"username"`
  Password string `json:"password"`
}

type JobPayload struct {
  Title string `json:"title"`
  Company string `json:"company"`
  Link string `json:"link"`
  Description string `json:"description"`
  TotalCompensation int `json:"total_compensation"`
}


func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
  var payload UserPayload
  
  err := json.NewDecoder(r.Body).Decode(&payload)
  if err != nil {
    log.Println(err)
    return
  }
  
  var user models.User

  user.Username = payload.Username 
  user.Password = payload.Password 

  err = app.models.DB.RegisterUser(user)
  if err != nil {
    app.logger.Println(err)
  }


  js, err := json.MarshalIndent(user, "", "\t")
  if err != nil {
    app.logger.Println(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(js)
}

func (app *application) insertJob(w http.ResponseWriter, r *http.Request) {
  var payload JobPayload
  
  err := json.NewDecoder(r.Body).Decode(&payload)
  if err != nil {
    log.Println(err)
    return
  }
  
  var job models.Job

  job.Title = payload.Title 
  job.Company = payload.Company 
  job.Link = payload.Link 
  job.Description = payload.Description 
  job.TotalCompensation = payload.TotalCompensation 
   
  err = app.models.DB.InsertJob(job)
  if err != nil {
    app.logger.Println(err)
  }

  js, err := json.MarshalIndent(job, "", "\t")
  if err != nil {
    app.logger.Println(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(js)

}

func (app *application) getAllUsers(w http.ResponseWriter, r *http.Request) {

}

func (app *application) getAllJobs(w http.ResponseWriter, r *http.Request) {

}





