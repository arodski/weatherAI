package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	OpenWeatherURL    string
	OpenWeatherAPIKey string
	Addr              string
}

func LoadConfig() *Config {
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current directory")
	}

	path := filepath.Join(curDir, ".env")

	err = godotenv.Load(path)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENWEATHER_API_KEY not set in .env file")
	}

	url := os.Getenv("OPENWEATHER_URL")
	if url == "" {
		log.Fatal("OPENWEATHER_URL not set in .env file")
	}

	addr := os.Getenv("ADDR")
	if addr == "" {
		log.Fatal("ADDR not set in .env file")
	}

	return &Config{
		OpenWeatherAPIKey: apiKey,
		OpenWeatherURL:    url,
		Addr:              addr,
	}
}
