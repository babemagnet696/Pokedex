package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please enter a caught pokemon to inspect")
	}
	pokemon := args[0]

	data, ok := cfg.pokedex[pokemon]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}
	stats := data.Stats
	types := data.Types

	fmt.Printf("Name: %s\n", data.Name)
	fmt.Printf("Height: %d\n", data.Height)
	fmt.Printf("Weight: %d\n", data.Weight)
	fmt.Println("Stats:")
	for _, statData := range stats {
		fmt.Printf("  -%s: %d\n", statData.Stat.Name, statData.BaseStat)
	}
	for _, typeData := range types {
		fmt.Printf("  - %s\n", typeData.Type.Name)
	}

	return nil

}