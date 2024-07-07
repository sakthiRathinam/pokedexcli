package main

import (
	"fmt"

	"github.com/sakthiRathinam/pokedexcli/internal/pokedex"
)


func commandGetPokemons(cfg *pokedex.PokedexConfig,pokedexElem string) error {
	pokedexResp,err := cfg.PokedexClient.GetPokemonOnLoc(cfg,pokedexElem)
	if err != nil {
		return err
	}
	_printPokemons(pokedexResp,pokedexElem)
	return nil
}

func _printPokemons(pokedexResp pokedex.PokedexSpecificLocationResponse,pokedexElem string) {
	fmt.Printf("Pokemon founded on this location %s\n",pokedexElem)
	for _,pokemonObjs := range pokedexResp.PokemonEncounters {
		fmt.Println(pokemonObjs.Pokemon.Name)
	}
}