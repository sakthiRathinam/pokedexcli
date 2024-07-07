package pokedex

import (
	"net/http"
	"time"

	"github.com/sakthiRathinam/pokedexcli/internal/pokedexcache"
)

type APICurrentState struct {
	PreviousURL *string
	NextURL *string
}

const BaseUrl = "https://pokeapi.co/api/v2/"

type PokedexClient struct {
	Client http.Client
}

type Pokemon struct {
	Name string
	Height int
	Weight int
	BaseExperience int
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

type PokedexConfig struct {
	PokedexClient PokedexClient
	Location APICurrentState
	PokedexCache pokedexcache.CacheStore
	PokemonStore map[string]Pokemon
}

func CreateClient() http.Client{
	return http.Client{Timeout: time.Minute}
}







