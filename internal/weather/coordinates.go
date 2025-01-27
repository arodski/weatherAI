package weather

import (
	"fmt"
	"net/http"
	"strconv"
)

type Coordinates struct {
	Lat float64
	Lon float64
}

func ParseCoordinates(r *http.Request) (*Coordinates, error) {
	latStr := r.URL.Query().Get("lat")
	lonStr := r.URL.Query().Get("lon")

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid latitude")
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid longitude")
	}

	return &Coordinates{Lat: lat, Lon: lon}, nil
}
