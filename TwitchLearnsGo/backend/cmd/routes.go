package main


import (
  "net/http"
  "github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
  router := httprouter.New()

  router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
  router.HandlerFunc(http.MethodGet, "/users/:id", app.getUser)
  router.HandlerFunc(http.MethodPost, "/register/user", app.registerUser)
  router.HandlerFunc(http.MethodPost, "/job/add", app.insertJob)
  router.HandlerFunc(http.MethodGet, "/jobs/:id", app.getJob)
  router.HandlerFunc(http.MethodGet, "/users/", app.getAllUsers)
  router.HandlerFunc(http.MethodGet, "/jobs/", app.getAllJobs)
  return router
}
