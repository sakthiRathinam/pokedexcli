package main

import (
	"github.com/sakthiRathinam/pokedexcli/internal/pokedex"
	"github.com/sakthiRathinam/pokedexcli/internal/pokedexcache"
)



func main() {
	pokedexClientObj := pokedex.PokedexClient{Client: pokedex.CreateClient()}
	pokedexResponseCache := pokedexcache.CreateCacheStore()
	
	pokedexConfig := pokedex.PokedexConfig{PokedexClient:pokedexClientObj,Location:pokedex.APICurrentState{PreviousURL: nil,NextURL: nil},PokedexCache: pokedexResponseCache}
	startRepl(&pokedexConfig)
}