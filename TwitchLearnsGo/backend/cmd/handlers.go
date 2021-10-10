package main

import (
	//"errors"
	"fmt"
	"backend/models"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
)

type AppStatus struct {
	Status string
}

type Message struct {
	Message string `json:"message"`
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
	Title             string `json:"title"`
	Company           string `json:"company"`
	Link              string `json:"link"`
	Description       string `json:"description"`
	TotalCompensation int    `json:"total_compensation"`
}

func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	var payload UserPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		return
	}

	var user models.User

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 12)
	if err != nil {
		app.logger.Println(err)
	}

	user.Username = payload.Username
	user.Password = string(hashPassword)

	err = app.models.DB.RegisterUser(user)
	if err != nil {
		app.logger.Println(err)
	}

	_message := Message{
		Message: "Successfully registered a user",
	}

	js, err := json.MarshalIndent(_message, "", "\t")
	if err != nil {
		app.logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func (app *application) signIn(w http.ResponseWriter, r *http.Request) {
	var payload UserPayload

	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.logger.Print("invalid id parameter")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		return
	}

	user, err := app.models.DB.GetUser(id)
	if err != nil {
		app.logger.Println("Could not get user")
	}

	fmt.Println(user.Password)

	hashPassword := user.Password	

	fmt.Println(hashPassword)
	fmt.Println(payload.Password)

	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(payload.Password))
  if err != nil {
		log.Println(err)
		_message := Message{
			Message: "Unauthorized",
		}

		js, err := json.MarshalIndent(_message, "", "\t")
		if err != nil {
			app.logger.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
		return

		}

	_message := Message{
		Message: "Successfully logged in",
	}

	js, err := json.MarshalIndent(_message, "", "\t")
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

	_message := Message{
		Message: "Successfully posted a job",
	}

	js, err := json.MarshalIndent(_message, "", "\t")
	if err != nil {
		app.logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)

}

func (app *application) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := app.models.DB.GetAllUsers()
	if err != nil {
		app.logger.Println("Could not get all users")
	}

	js, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		app.logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func (app *application) getAllJobs(w http.ResponseWriter, r *http.Request) {
	jobs, err := app.models.DB.GetAllJobs()
	if err != nil {
		app.logger.Println("Could not get all users")
	}

	js, err := json.MarshalIndent(jobs, "", "\t")
	if err != nil {
		app.logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

type TestPayload struct {
	One string `json:"one"`
	Two string `json:"two"`
	Three string `json:"three"`
} 

func (app *application) insertForm(w http.ResponseWriter, r *http.Request) {
	var payload TestPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		return
	}

	var formload models.Payload

	formload.One = payload.One
	formload.Two = payload.Two
	formload.Three = payload.Three

	js, err := json.MarshalIndent(formload, "", "\t")
	if err != nil {
		app.logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)

}




















