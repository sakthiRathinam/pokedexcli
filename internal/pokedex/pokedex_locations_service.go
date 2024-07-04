package pokedex

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/sakthiRathinam/pokedexcli/internal/pokedexcache"
)
type PokedexLocationsResponse struct {
	Count    *int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Locations  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (pokedexClient *PokedexClient) GetLocationsNext(cfg *PokedexConfig) (PokedexLocationsResponse,error) {
	locationUrl := BaseUrl + "location-area"
	 
	if cfg.Location.NextURL != nil{
		locationUrl = *cfg.Location.NextURL
	}
	return _getLocations(locationUrl,cfg)
	
}




func (pokedexClient *PokedexClient) GetLocationsPrevious(cfg *PokedexConfig) (PokedexLocationsResponse,error) {
	pokedexResponse := PokedexLocationsResponse{}
	if cfg.Location.PreviousURL != nil {
		return _getLocations(*cfg.Location.PreviousURL,cfg)
	}
	return pokedexResponse,errors.New("no previous location available")
}

func _getLocations(locationURL string,cfg *PokedexConfig) (PokedexLocationsResponse,error){
	cacheEntry,err := cfg.PokedexCache.GetCacheResponse(locationURL)
	if err != nil {
		return _getLocationsHittingPokeAPI(locationURL,&cfg.PokedexCache)
	}
	return _convertBytesToLocationResp(cacheEntry.Val)
}


func _getLocationsHittingPokeAPI(locationURL string,pokedexCache *pokedexcache.CacheStore) (PokedexLocationsResponse,error) {
	
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

	locationResp,err := _convertBytesToLocationResp(responseBody)
	if err != nil {
		return locationResp,errors.New("error occured while parsing json")
	}
	stored := pokedexCache.StoreCacheEntry(locationURL,responseBody,10)
	if stored != nil {
		fmt.Println("storing the cache failed")
	}
	return locationResp,nil

}


func _convertBytesToLocationResp(responseBody []byte) (PokedexLocationsResponse,error){
	var locationsResp PokedexLocationsResponse
	jsonMarshal := json.Unmarshal(responseBody,&locationsResp)
	if jsonMarshal != nil {
		return locationsResp,errors.New("error occured while unmarshal location")
	}
	return locationsResp,nil
}