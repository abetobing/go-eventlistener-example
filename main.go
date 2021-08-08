package main

import (
	"context"
	"net/http"

	"github.com/abetobing/go-eventlistener-example/events"
	"github.com/abetobing/go-eventlistener-example/handler"
	"github.com/abetobing/go-eventlistener-example/listeners"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	userEvent *events.UserEvent = &events.UserEvent{}
)

func main() {
	r := chi.NewRouter()

	userEvent.Subscribe("user_created", listeners.UserCreated)
	userEvent.Subscribe("user_created", listeners.SlackNotifier)

	r.Use(middleware.Logger)
	r.Route("/users", func(r chi.Router) {
		r.Use(EventMiddleware)
		r.Post("/", handler.UserCreateHandler)
	})
	http.ListenAndServe(":3000", r)
}

func EventMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user_event_ctx_key", userEvent)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
