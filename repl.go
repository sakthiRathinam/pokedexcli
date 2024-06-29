package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
type cliCommand struct {
	name string
	description string
	callback func() error
}

func cleanInput(input string) string {
	removeSpace := strings.TrimSpace(input)
	cleanedInput := strings.ToLower(removeSpace)
	return cleanedInput
}

func startRepl(){
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
        callback:    clear,
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