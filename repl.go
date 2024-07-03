package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sakthiRathinam/pokedexcli/internal/pokedex"
	"github.com/sakthiRathinam/pokedexcli/internal/pokedexcache"
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

func displayLenOfCache(cfg *pokedex.PokedexConfig) error{
	fmt.Println(len(cfg.PokedexCache.Store))
	return nil
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
	"mapb": {
        name:        "Locations",
        description: "Get the previous on the pokedex world",
        callback:    command_mapb,
    },
	"cachel": {
        name:        "Cache Lenght",
        description: "display the length of the cache",
        callback:    displayLenOfCache,
    }, // only for debugging
	
	}
	reader := bufio.NewScanner(os.Stdin)
	closeChan := make(chan int)
	go pokedexcache.ReapLoop(&cfg.PokedexCache,&closeChan,3)
	fmt.Printf("pokedoxcli> ")
	for reader.Scan(){
		text := cleanInput(reader.Text())
		cliCommand, exists := commandsMap[text]
		if exists{
			err := cliCommand.callback(cfg)
			if err != nil{
				fmt.Println(err)
			}
		}
		if text == "exit"{
			closeChan <- 1
			break
		}
		fmt.Printf("pokedoxcli> ")
	}

}