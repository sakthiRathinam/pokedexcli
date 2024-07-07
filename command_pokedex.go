package main

import (
	"fmt"

	"github.com/sakthiRathinam/pokedexcli/internal/pokedex"
)



func commandPokedex(cfg *pokedex.PokedexConfig,pokedexElem string) error {
	if len(cfg.PokemonStore) == 0 {
		fmt.Println("Your pokedex collection is empty")
		return nil
	}
	fmt.Println("Your pokedex:")
	for pokemon := range cfg.PokemonStore{
		fmt.Printf("    - %s\n",pokemon)
	}
	return nil
}