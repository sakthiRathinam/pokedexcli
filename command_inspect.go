package main

import (
	"fmt"

	"github.com/sakthiRathinam/pokedexcli/internal/pokedex"
)


func commandInspect(cfg *pokedex.PokedexConfig,pokedexElem string) error {
	pokemon, ok := cfg.PokemonStore[pokedexElem]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	_printPokemonDetails(pokemon)
	return nil 

}


func _printPokemonDetails(pokemon pokedex.Pokemon){
	fmt.Printf("Name: %s\n",pokemon.Name)
	fmt.Printf("Height: %d\n",pokemon.Height)
	fmt.Printf("Weight: %d\n",pokemon.Weight)
	fmt.Printf("Experience: %d\n",pokemon.BaseExperience)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("    -%s: %d\n",stat.Stat.Name,stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("    -%s\n",pokemonType.Type.Name)
	}
	
}

