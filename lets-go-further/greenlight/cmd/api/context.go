package main

import (
	"context"
	"github.com/amokstakov/greenlight/internal/data"
	"net/http"
)

// Define a custom contextKey, with the underlying type string
type contextKey string

// convert the string user to a contextKey type
// Use this constant as the key for getting and setting user info
const userContextKey = contextKey("user")

// returns a new copy of the request with provided User struct added to the context
func (app *application) contextSetUser(r *http.Request, user *data.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

// we will only need this when we expect a user struct value in the context
func (app *application) contextGetUser(r *http.Request) *data.User {
	user, ok := r.Context().Value(userContextKey).(*data.User)
	if !ok {
		panic("missing user value in request context")
	}

	return user
}
