package go_geonames

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	Username string
}

type SearchRequest struct {
	Latitude  float64
	Longitude float64
	Radius    float64
}

type SearchFilter string

const (
	ServiceBaseURL                                  = "http://api.geonames.org"
	CityWithPopulationGreaterThan1000  SearchFilter = "cities1000"
	CityWithPopulationGreaterThan5000  SearchFilter = "cities5000"
	CityWithPopulationGreaterThan15000 SearchFilter = "cities15000"
)

func (c *Client) GetNearbyCities(req *SearchRequest, filter SearchFilter) ([]City, error) {
	rawURL, err := url.JoinPath(ServiceBaseURL, "findNearbyPlaceNameJSON")
	if err != nil {
		return nil, err
	}

	URL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	params := URL.Query()
	params.Add("username", c.Username)
	params.Add("lat", fmt.Sprintf("%.4f", req.Latitude))
	params.Add("lng", fmt.Sprintf("%.4f", req.Longitude))
	params.Add("radius", fmt.Sprintf("%.2f", req.Radius))
	params.Add("cities", string(filter))

	URL.RawQuery = params.Encode()
	finalURL := URL.String()
	response, err := http.Get(finalURL)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	cities := &Geonames{Cities: make([]City, 0)}
	if err = json.Unmarshal(body, cities); err != nil {
		return nil, err
	}

	return cities.Cities, nil
}

type City struct {
	ID         int64  `json:"geonameId"`
	Name       string `json:"name"`
	Latitude   string `json:"lat"`
	Longitude  string `json:"lng"`
	Population int64  `json:"population"`
	AdminArea1 string `json:"adminCode1"`
	Country    string `json:"countryName"`
}

type Geonames struct {
	Cities []City `json:"geonames"`
}
