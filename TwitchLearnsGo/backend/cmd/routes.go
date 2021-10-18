package main


import (
  "net/http"
  "github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
  router := httprouter.New()
  // Add our custom error handlers 
  router.NotFound = http.HandlerFunc(app.notFoundResponse)
  router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

  router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)


  // Everythin below is going to be a future verson implementin
  router.HandlerFunc(http.MethodGet, "/users/:id", app.getUser)
  router.HandlerFunc(http.MethodPost, "/register/user", app.registerUser)
  router.HandlerFunc(http.MethodPost, "/job/add", app.insertJob)
  router.HandlerFunc(http.MethodGet, "/jobs/:id", app.getJob)
  router.HandlerFunc(http.MethodGet, "/users/", app.getAllUsers)
  router.HandlerFunc(http.MethodGet, "/jobs/", app.getAllUsers)
  router.HandlerFunc(http.MethodPost, "/login/:id", app.signIn)
  router.HandlerFunc(http.MethodPost, "/form/", app.insertForm)
  return app.EnableCors(router)
}
