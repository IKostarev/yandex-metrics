package app

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"yandex-metrics/internal/config"
	"yandex-metrics/internal/handlers"
)

func Run() {
	route := chi.NewRouter()
	route.Route("/", func(r chi.Router) {
		r.Route("/update", func(r chi.Router) {
			r.Post("/gauge/{name}/{value}", handlers.GaugeHandler)
			r.Post("/counter/{name}/{value}", handlers.CounterHandler)
			r.Post("/{all}/{name}/{value}", handlers.BadRequestHandler)
		})
	})

	if err := http.ListenAndServe(config.ADDRESS, route); err != nil {
		log.Fatalf("Error starting server on port %s...\n", config.ADDRESS)
	}
}
