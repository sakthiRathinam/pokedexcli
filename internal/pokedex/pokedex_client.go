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

func (pokedexClient *PokedexClient) GetLocationsNext(cfg *PokedexConfig,next bool) (PokedexLocationsResponse,error) {
	locationUrl := BaseUrl + "location"
	var locationsResp PokedexLocationsResponse
	 
	response,err := http.Get(locationUrl)
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



