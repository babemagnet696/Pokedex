package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/babemagnet696/Pokedex/internal/pokeapi"
)


type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println()
		fmt.Print("Pokedex > ")
		reader.Scan()
		fmt.Println()


		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName, args := words[0], words[1:]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}	
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:     commandMapf,
		},
		"mapb": {
			name:        "mapb (map back)",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Shows all Pokemon in a selected area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon!",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon in your pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught pokemon from your pokedex",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	cleanText := strings.Fields(lowerText)
	return cleanText
}

