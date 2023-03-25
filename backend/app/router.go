package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewRouter ...
func NewRouter() chi.Router {

	//Create main router
	r := chi.NewRouter()
	// r.Use(mid)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	/**
	 * Routes
	 */
	r.Get("/users/login", Login)
	// r.Methods("GET").Path("/users/login/").HandlerFunc(Login)

	return r
}
