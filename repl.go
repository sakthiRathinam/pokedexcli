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
	callback func(*pokedex.PokedexConfig, string) error
}

func cleanInput(input string) (string,string) {
	splitBySpace := strings.Split(input, " ")
	for index,elem := range splitBySpace{
		removeSpace := strings.TrimSpace(elem)
		cleanedInput := strings.ToLower(removeSpace)
		splitBySpace[index] = cleanedInput
	}
	if len(splitBySpace) < 2 {
		return splitBySpace[0],""
	}
	return splitBySpace[0],splitBySpace[1]
}

func displayLenOfCache(cfg *pokedex.PokedexConfig, pokedexElem string) error{
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
	"explore": {
        name:        "Explore pokemons",
        description: "explore pokemons for that particular location",
        callback:    command_get_pokemons,
    }, 
	"catch": {
        name:        "Explore pokemons",
        description: "explore pokemons for that particular location",
        callback:    command_catch,
    }, 
	"inspect": {
        name:        "Inspect pokemon",
        description: "Inspect the pokemon and caught or not, If caught then display the details",
        callback:    command_inspect,
    }, 
	"pokedex": {
        name:        "Pokemons",
        description: "Display all the pokemons in our bucket",
        callback:    command_pokedex,
    }, 
	
	}
	
	reader := bufio.NewScanner(os.Stdin)
	closeChan := make(chan int)
	go pokedexcache.ReapLoop(&cfg.PokedexCache,&closeChan,3)
	fmt.Printf("pokedoxcli> ")
	for reader.Scan(){
		text,pokedexElem := cleanInput(reader.Text())
		
		cliCommand, exists := commandsMap[text]
		if exists{
			err := cliCommand.callback(cfg,pokedexElem)
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



