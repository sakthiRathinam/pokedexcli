package main

import (
	"fmt"

	"github.com/sakthiRathinam/pokedexcli/internal/pokedex"
)


func commandRelease(cfg *pokedex.PokedexConfig,pokedexElem string) error {
	_, exists := cfg.PokemonStore[pokedexElem]
	if !exists {
		fmt.Printf("%s is not been caught yet to release it\n",pokedexElem)
		return nil
	}
	delete(cfg.PokemonStore,pokedexElem)
	fmt.Printf("%s has released to the pokemon world and saying goodbyeeeeee.....\n",pokedexElem)
	return nil
}