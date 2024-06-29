package main

import (
	"errors"
	"fmt"

	"github.com/sakthiRathinam/pokedexcli/internal/pokedex"
)

func command_map(cfg *pokedex.PokedexConfig) error {
	locationsResp,err := cfg.PokedexClient.GetLocations()
	if err != nil {
		return errors.New("error occured while getting the locations")
	}
	fmt.Println("Pokedex world locations")
	if locationsResp.Next != nil {
		cfg.Location.NextURL = locationsResp.Next
	}
	
	if locationsResp.Previous != nil {
		cfg.Location.PreviousURL = locationsResp.Previous
	}
	for _,location := range locationsResp.Locations {
		fmt.Println(location.Name)
	}
	return nil
}