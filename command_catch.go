package main

import (
	"fmt"
	"time"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("No pokemon to catch entered")
	}

	pokemon := args[0]

	pokemonResp, err := cfg.pokeapiClient.PokemonDetails(pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	time.Sleep(2 * time.Second)
	chance := rand.Intn(500)
	if chance >= pokemonResp.BaseExperience {
		cfg.pokedex[pokemon] = pokemonResp
		fmt.Printf("%s was caught!\n", pokemon)
		fmt.Println("You may now inspect it with the inspect command.")
		return nil
	}

	fmt.Printf("%s escaped!\n", pokemon)
	return nil
}