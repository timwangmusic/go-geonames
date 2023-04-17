package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	username string
}

type SearchRequest struct {
	latitude  float64
	longitude float64
	radius    float64
}

type SearchFilter string

const (
	ServiceBaseURL                                  = "http://api.geonames.org"
	CityWithPopulationGreaterThan1000  SearchFilter = "cities1000"
	CityWithPopulationGreaterThan5000  SearchFilter = "cities5000"
	CityWithPopulationGreaterThan15000 SearchFilter = "cities15000"
)

func (c *Client) getNearbyCities(req *SearchRequest, filter SearchFilter) ([]City, error) {
	rawURL, err := url.JoinPath(ServiceBaseURL, "findNearbyPlaceNameJSON")
	if err != nil {
		return nil, err
	}

	URL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	params := URL.Query()
	params.Add("username", c.username)
	params.Add("lat", fmt.Sprintf("%.4f", req.latitude))
	params.Add("lng", fmt.Sprintf("%.4f", req.longitude))
	params.Add("radius", fmt.Sprintf("%.2f", req.radius))
	params.Add("cities", string(filter))

	URL.RawQuery = params.Encode()
	response, err := http.Get(URL.String())
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
