package handlers 

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/cankurttekin/gotutorial/11-api/internal/middleware"
)

func Handler(r *chi.Mux) {
	// global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		// middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}

