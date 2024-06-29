package main

import (
	"errors"
	"fmt"
)
func CommandExit() error{
	fmt.Printf("pokedexcli is shutting down")
	return errors.New("Quit!!!!!!")
}