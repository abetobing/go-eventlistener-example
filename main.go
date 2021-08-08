package main

import (
	"net/http"

	"github.com/abetobing/go-eventlistener-example/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Route("/users", func(r chi.Router) {
		r.Post("/", handler.UserCreateHandler)
	})
	http.ListenAndServe(":3000", r)
}
