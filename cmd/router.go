package main

import (
	"github.com/aadi-1024/quickshare/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

// NewRouter sets up and returns a router
func NewRouter() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)

	mux.Get("/", handlers.Home)
	mux.Post("/verify", handlers.Verify)

	return mux
}
