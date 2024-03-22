package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/jignyasamishra/Go-API/internal/middleware"
)

func Handler(r *chi.Mux) {
	//GLOBAL MIDDLERWARE
	r.Use(chimiddle.StringSplashes)
	r.Route("/account", func(router chi.Router) {
		//middleware for /account route
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}
