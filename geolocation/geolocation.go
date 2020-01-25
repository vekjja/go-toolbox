package geolocation

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Data : Data returned from opendatasoft
type Data struct {
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

// FromIP : get location data based on IP
func FromIP() (Data, error) {

	locationData := Data{}

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

// Locate : Use Google Location Servicies to get location data from search string
func Locate(location string) (Data, error) {

	locationData := Data{}
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
