package main

import "fmt"

func commandHelp() error {
	commandsList := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range commandsList {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil

}