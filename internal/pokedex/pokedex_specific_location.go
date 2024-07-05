package pokedex

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/sakthiRathinam/pokedexcli/internal/pokedexcache"
)


type PokedexSpecificLocationResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}


func (pokedexClient *PokedexClient) GetPokemonOnLoc(cfg *PokedexConfig,location string) (PokedexSpecificLocationResponse,error) {
	pokedexURL := BaseUrl + "location-area/" + location + "/"
	return _getPokeOnLoc(pokedexURL,cfg)
}

func _getPokeOnLoc(locationURL string,cfg *PokedexConfig) (PokedexSpecificLocationResponse,error){
	cacheEntry,err := cfg.PokedexCache.GetCacheResponse(locationURL)
	if err != nil {
		return _getPokeOnLocByHittingPokeAPI(&cfg.PokedexCache,locationURL)
	}
	return _convertBytesToPokemonResp(cacheEntry.Val)
}


func _getPokeOnLocByHittingPokeAPI(pokedexCache *pokedexcache.CacheStore,locationURL string) (PokedexSpecificLocationResponse,error) {
	var locationResp PokedexSpecificLocationResponse;
	response,err := http.Get(locationURL)
	if err != nil {
		return locationResp,errors.New("error occured while fetching location")
	}

	responseBody,err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return locationResp,errors.New("error occured while parsing location")
	}
	pokedexCache.StoreCacheEntry(locationURL,responseBody,20)
	jsonMarshal := json.Unmarshal(responseBody,&locationResp)
	if jsonMarshal != nil {
		return locationResp,errors.New("error occured while unmarshal location")
	}
	return locationResp,nil
}

func _convertBytesToPokemonResp(responseBody []byte) (PokedexSpecificLocationResponse,error){
	var locationsResp PokedexSpecificLocationResponse
	jsonMarshal := json.Unmarshal(responseBody,&locationsResp)
	if jsonMarshal != nil {
		return locationsResp,errors.New("error occured while unmarshal location")
	}
	return locationsResp,nil
}


