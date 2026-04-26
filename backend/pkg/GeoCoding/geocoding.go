package geocoding

import (
	"fmt"
	"net/http"

	"github.com/bytedance/sonic"
)

type GeoCoder struct {
}

const (
	geocoderBaseAPIUrl  = "https://geocode.gate.petersburg.ru"
	parseStreetEndpoint = "/parse/eas"
)

func New() *GeoCoder {
	return &GeoCoder{}
}

func (gc *GeoCoder) GetCoords(address string) ([]float64, error) {
	url := geocoderBaseAPIUrl + parseStreetEndpoint
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error getting coordinates: %w", err)
	}

	q := req.URL.Query()
	q.Add("street", address)

	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error getting coordinates: %w", err)
	}
	defer resp.Body.Close()

	result := new(parseAddressResponse)
	err = sonic.ConfigDefault.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, fmt.Errorf("Error getting coordinates: %w", err)
	}
	return []float64{result.Latitude, result.Longitude}, nil
}

type parseAddressResponse struct {
	ID         int     `json:"ID"`
	BuildingID int     `json:"Building_ID"`
	Name       string  `json:"Name"`
	Longitude  float64 `json:"Longitude"`
	Latitude   float64 `json:"Latitude"`
}
