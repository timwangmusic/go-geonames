package main

import (
	"fmt"
	"os"
)

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

func main() {
	fmt.Println("welcome to use golang client for geonames")

	username := os.Getenv("username")
	if username == "" {
		username = "demo"
	}
	c := Client{username: username}

	if cities, err := c.getNearbyCities(&SearchRequest{
		latitude:  37.4419,
		longitude: -122.1430,
		radius:    20.0,
	}, CityWithPopulationGreaterThan5000); err != nil {
		fmt.Println(err.Error())
	} else {
		for _, city := range cities {
			fmt.Println(city.Name)
		}
	}
}
