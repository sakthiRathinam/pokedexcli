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

type PokedexConfig struct {
	PokedexClient PokedexClient
	Location APICurrentState
	PokedexCache pokedexcache.CacheStore
}

func CreateClient() http.Client{
	return http.Client{Timeout: time.Minute}
}







