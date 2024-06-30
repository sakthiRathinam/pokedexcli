package main

import (
	"fmt"

	"github.com/sakthiRathinam/pokedexcli/internal/pokedex"
)
func CommandExit(cfg *pokedex.PokedexConfig) error{
	fmt.Printf("pokedexcli is shutting down")
	return nil
}