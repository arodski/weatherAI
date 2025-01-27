package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Service interface {
	GetWeather(lat, lon float64) (*WeatherData, error)
}

type OpenWeatherService struct {
	url    string
	apiKey string
}

func NewOpenWeatherService(url string, apiKey string) *OpenWeatherService {
	return &OpenWeatherService{
		url:    url,
		apiKey: apiKey,
	}
}

func (s *OpenWeatherService) GetWeather(lat, lon float64) (*WeatherData, error) {
	url := fmt.Sprintf("%s?lat=%f&lon=%f&appid=%s&units=imperial", s.url, lat, lon, s.apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result OpenWeatherResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	temp := Temperature{Value: result.Main.Temp}
	classification := temp.Classify()

	return &WeatherData{
		Condition:   result.Weather[0].Main,
		Temperature: classification,
	}, nil
}
