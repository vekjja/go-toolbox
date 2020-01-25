package location

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GeoLocationData : Data returned from opendatasoft
type GeoLocationData struct {
	City        string  `json:"city"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	Region      string  `json:"region"`
	RegionCode  string  `json:"region_code"`
	PostalCode  string  `json:"postal_code"`
	Country     string  `json:"country"`
	CountryCode string  `json:"country_code"`
	Timezone    string  `json:"timezone"`
	IP          string  `json:"ip"`
}

// both of the folloing APIs are not provided by me 😅

// GeoLocateFromIP : get location data based on IP
func GeoLocateFromIP() (GeoLocationData, error) {

	locationData := GeoLocationData{}

	url := "https://telize.j3ss.co/geoip"
	res, err := http.Get(url)
	if err != nil {
		return locationData, err
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return locationData, err
	}

	json.Unmarshal(responseData, &locationData)

	return locationData, nil

}

// GeoLocate : Use Google Location Servicies to get location data from search string
func GeoLocate(location string) (GeoLocationData, error) {

	locationData := GeoLocationData{}
	url := "https://geocode.jessfraz.com/geocode"

	reqBody, _ := json.Marshal(map[string]string{
		"Location": location,
	})

	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return locationData, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	json.Unmarshal(body, &locationData)

	return locationData, nil

}
