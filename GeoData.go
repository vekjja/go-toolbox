package toolbox

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type NominatimResponse []struct {
	Address struct {
		ISO31662Lvl4  string `json:"ISO3166-2-lvl4"`
		Borough       string `json:"borough"`
		City          string `json:"city"`
		Country       string `json:"country"`
		CountryCode   string `json:"country_code"`
		Neighbourhood string `json:"neighbourhood"`
		Postcode      string `json:"postcode"`
		Road          string `json:"road"`
		Shop          string `json:"shop"`
		Suburb        string `json:"suburb"`
	} `json:"address"`
	Addresstype string   `json:"addresstype"`
	Boundingbox []string `json:"boundingbox"`
	Category    string   `json:"category"`
	DisplayName string   `json:"display_name"`
	Importance  float64  `json:"importance"`
	Lat         string   `json:"lat"`
	Licence     string   `json:"licence"`
	Lon         string   `json:"lon"`
	Name        string   `json:"name"`
	OsmID       int      `json:"osm_id"`
	OsmType     string   `json:"osm_type"`
	PlaceID     int      `json:"place_id"`
	PlaceRank   int      `json:"place_rank"`
	Type        string   `json:"type"`
}

// GetGeoData: Get geographical data from Nominatim
func GetGeoData(location string) (NominatimResponse, error) {
	baseURL := "https://nominatim.openstreetmap.org/search"
	encoded := url.QueryEscape(location)
	reqURL := fmt.Sprintf("%s?q=%s&format=json&limit=1", baseURL, encoded)

	client := &http.Client{}
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get geo data: %s", resp.Status)
	}

	var data NominatimResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("no data found for location: %s", location)
	}

	return data, nil
}
