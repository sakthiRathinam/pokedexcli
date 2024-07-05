package main

import (
	"fmt"

	"github.com/sakthiRathinam/pokedexcli/internal/pokedex"
)

func command_map(cfg *pokedex.PokedexConfig,pokedexElem string) error {
	locationsResp,err := cfg.PokedexClient.GetLocationsNext(cfg)
	if err != nil {
		return err
	}

	cfg.Location.NextURL = locationsResp.Next
	cfg.Location.PreviousURL = locationsResp.Previous
	
	_printLocations(locationsResp)
	return nil
}


func command_mapb(cfg *pokedex.PokedexConfig,pokedexElem string) error {
	locationsResp,err := cfg.PokedexClient.GetLocationsPrevious(cfg)
	if err != nil {
		return err
	}

	cfg.Location.NextURL = locationsResp.Next
	cfg.Location.PreviousURL = locationsResp.Previous
    
	_printLocations(locationsResp)
	return nil
}


func _printLocations(locationsResp pokedex.PokedexLocationsResponse) {
	fmt.Println("Pokedex world locations")
	for _,location := range locationsResp.Locations {
		fmt.Println(location.Name)
	}
}