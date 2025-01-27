package main

import (
	"log"
	"net/http"

	"github.com/arodski/weatherAI/internal/config"
	"github.com/arodski/weatherAI/internal/weather"
)

type application struct {
	config         config.Config
	weatherService weather.Service
}

func (app *application) mount() http.Handler {
	m := *http.NewServeMux()

	m.Handle("GET /v1/weather", http.HandlerFunc(app.getWeatherHandler))

	return &m
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:    app.config.Addr,
		Handler: mux,
	}

	log.Printf("server has started at %s", app.config.Addr)

	return srv.ListenAndServe()
}
