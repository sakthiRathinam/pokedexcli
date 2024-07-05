package main

import (
	"fmt"

	"github.com/sakthiRathinam/pokedexcli/internal/pokedex"
)
func CommandExit(cfg *pokedex.PokedexConfig,pokedexElem string) error{
	fmt.Println("pokedexcli is shutting down")
	return nil
}