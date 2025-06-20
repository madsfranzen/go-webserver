package main

import (
	"net/http"

	"github.com/madsfranzen/go-webserver/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func setupRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", handlers.GetUsers)
		r.Post("/", handlers.CreateUser)
	})
	return r
}
