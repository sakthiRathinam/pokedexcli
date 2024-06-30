package pokedex

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type APICurrentState struct {
	PreviousURL *string
	NextURL *string
}


type PokedexLocationsResponse struct {
	Count    *int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Locations  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const BaseUrl = "https://pokeapi.co/api/v2/"

type PokedexClient struct {
	Client http.Client
}

type PokedexConfig struct {
	PokedexClient PokedexClient
	Location APICurrentState
}

func CreateClient() http.Client{
	return http.Client{Timeout: time.Minute}
}

func (pokedexClient *PokedexClient) GetLocationsNext(cfg *PokedexConfig) (PokedexLocationsResponse,error) {
	locationUrl := BaseUrl + "location"
	 
	if cfg.Location.NextURL != nil{
		locationUrl = *cfg.Location.NextURL
	}

	return _getLocations(locationUrl)
}


func (pokedexClient *PokedexClient) GetLocationsPrevious(cfg *PokedexConfig) (PokedexLocationsResponse,error) {
	pokedexResponse := PokedexLocationsResponse{}
	if cfg.Location.PreviousURL != nil {
		return _getLocations(*cfg.Location.PreviousURL)
	}
	return pokedexResponse,errors.New("no previous location available")
}

func _getLocations(locationURL string) (PokedexLocationsResponse,error) {
	var locationsResp PokedexLocationsResponse

	response,err := http.Get(locationURL)
	if err != nil {
		return locationsResp,errors.New("error occured while fetching location")
	}

	responseBody,err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return locationsResp,errors.New("error occured while parsing location")
	}

	jsonMarshal := json.Unmarshal(responseBody,&locationsResp)
	if jsonMarshal != nil {
		fmt.Println(err)
		return locationsResp,errors.New("error occured while unmarshal location")
	}

	return locationsResp,nil

}



