package main

import (
	"fmt"

	"github.com/sakthiRathinam/pokedexcli/internal/pokedex"
)
func displayHelpText(cfg *pokedex.PokedexConfig,pokedexElem string) error{
	fmt.Printf(
        "Welcome to pokedexcli! These are the available commands: \n",
    )
    fmt.Println("help    - Show available commands")
    fmt.Println("clear   - Clear the terminal screen")
    fmt.Println("exit    - Closes the terminal screen")
    fmt.Println("map     - Shows the locations on the pokemon world")
    fmt.Println("mapb    - Goback to previous locations on the pokemon world")
	return nil
}


