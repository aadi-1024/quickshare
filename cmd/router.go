package main

import (
	"fmt"
	"github.com/aadi-1024/quickshare/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

// NewRouter sets up and returns a router
func NewRouter(hash string) http.Handler {
	mux := chi.NewRouter()

	if !app.InProd {
		mux.Use(middleware.Logger)
	}

	mux.Get("/", handlers.Home)
	mux.Post("/verify", handlers.Verify)

	mux.Get(fmt.Sprintf("/%v", hash), handlers.File)
	return mux
}
