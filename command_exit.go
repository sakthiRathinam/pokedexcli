package main

import (
	"fmt"

	"github.com/sakthiRathinam/pokedexcli/internal/pokedex"
)
func commandExit(cfg *pokedex.PokedexConfig,pokedexElem string) error{
	fmt.Println("pokedexcli is shutting down")
	return nil
}