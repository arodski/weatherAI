package main

import (
	"encoding/json"
	"net/http"

	"github.com/arodski/weatherAI/internal/weather"
)

func (app *application) getWeatherHandler(w http.ResponseWriter, r *http.Request) {
	coords, err := weather.ParseCoordinates(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	weatherData, err := app.weatherService.GetWeather(coords.Lat, coords.Lon)
	if err != nil {
		http.Error(w, "Error fetching weather data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherData)
}
