package main

import "github.com/sakthiRathinam/pokedexcli/internal/pokedex"



func main() {
	pokedexClientObj := pokedex.PokedexClient{Client: pokedex.CreateClient()}
	pokedexConfig := pokedex.PokedexConfig{PokedexClient:pokedexClientObj,Location:pokedex.APICurrentState{PreviousURL: nil,NextURL: nil}}
	startRepl(&pokedexConfig)
}