package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)


type cliCommand struct {
	name string
	description string
	callback func() error
}

func displayHelpText() error{
	fmt.Printf(
        "Welcome to pokedexcli! These are the available commands: \n",
    )
    fmt.Println(".help    - Show available commands")
    fmt.Println(".clear   - Clear the terminal screen")
    fmt.Println(".quit    - Closes the terminal screen")
	return nil
}

func commandExit() error{
	fmt.Printf("pokedexcli is shutting down")
	return errors.New("Quit!!!!!!")
}



func cleanInput(input string) string {
	removeSpace := strings.TrimSpace(input)
	cleanedInput := strings.ToLower(removeSpace)
	return cleanedInput
}

func main() {

	commandsMap := map[string]cliCommand{
	"help": {
		name:"Help Message",
		description:"List all the commands with the description",
		callback: displayHelpText,
	},
	"exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
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