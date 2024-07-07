package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sakthiRathinam/pokedexcli/internal/pokedex"
)


func command_catch(cfg *pokedex.PokedexConfig,pokedexElem string) error {
	pokedexResp,err := cfg.PokedexClient.GetPokemonDetails(cfg,pokedexElem)
	if err != nil {
		return err
	}
	_tryCatchingPokemon(cfg,pokedexResp)
	return nil
}


func _tryCatchingPokemon(cfg *pokedex.PokedexConfig, pokedexResp pokedex.PokemonDetailsResp) error {
	
	max_val := _getMaxValBasedOnExperience(pokedexResp.BaseExperience)
	random_num_to_catch := rand.Intn(max_val) 
	user_num_to_catch := rand.Intn(max_val)
	fmt.Println("Catching the pokemon....")
	time.Sleep(2 * time.Second)
	if random_num_to_catch == user_num_to_catch {
		fmt.Printf("Wohooo you caught the pokemon %s\n",pokedexResp.Name)
		pokemonObj := _getPokemonObj(pokedexResp)
		cfg.PokemonStore[pokedexResp.Name] = pokemonObj
		return nil
	}
	fmt.Println("Pokemon escaped...try catching again")
	return nil
}


func _getMaxValBasedOnExperience(baseExperience int) int {
	return int(baseExperience / 20)
}


func _getPokemonObj(pokedexResp pokedex.PokemonDetailsResp)pokedex.Pokemon {
	return pokedex.Pokemon{
		Name:pokedexResp.Name,
		Height:pokedexResp.Height,
		BaseExperience:pokedexResp.BaseExperience,
		Weight:pokedexResp.Weight,
		Stats: pokedexResp.Stats,
		Types: pokedexResp.Types,
	}
}


