package main

import "fmt"
func DisplayHelpText() error{
	fmt.Printf(
        "Welcome to pokedexcli! These are the available commands: \n",
    )
    fmt.Println(".help    - Show available commands")
    fmt.Println(".clear   - Clear the terminal screen")
    fmt.Println(".exit    - Closes the terminal screen")
	return nil
}