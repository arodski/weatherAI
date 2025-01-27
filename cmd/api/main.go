package main

import (
	"log"

	"github.com/arodski/weatherAI/internal/config"
	"github.com/arodski/weatherAI/internal/weather"
)

func main() {
	cfg := config.LoadConfig()

	weatherService := weather.NewOpenWeatherService(cfg.OpenWeatherURL, cfg.OpenWeatherAPIKey)

	app := &application{
		config:         *cfg,
		weatherService: weatherService,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
