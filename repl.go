package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sakthiRathinam/pokedexcli/internal/pokedex"
)
type cliCommand struct {
	name string
	description string
	callback func(*pokedex.PokedexConfig) error
}

func cleanInput(input string) string {
	removeSpace := strings.TrimSpace(input)
	cleanedInput := strings.ToLower(removeSpace)
	return cleanedInput
}

func startRepl(cfg *pokedex.PokedexConfig){
commandsMap := map[string]cliCommand{
	"help": {
		name:"Help Message",
		description:"List all the commands with the description",
		callback: DisplayHelpText,
	},
	"exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    CommandExit,
    },
	"clear": {
        name:        "Clear",
        description: "Clear the current screen",
        callback:    clearScreen,
    },
	"map": {
        name:        "Locations",
        description: "Get the locations on the pokedex world",
        callback:    command_map,
    },
	
	}
	reader := bufio.NewScanner(os.Stdin)
	fmt.Printf("pokedoxcli> ")
	for reader.Scan(){
		text := cleanInput(reader.Text())
		cliCommand, exists := commandsMap[text]
		if exists{
			err := cliCommand.callback()
			if err != nil{
				break
			}
		}
		fmt.Printf("pokedoxcli> ")
	}

}