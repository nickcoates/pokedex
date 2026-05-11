package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location area name")
	}
	fmt.Printf("Exploring %v...\n", args[0])

	exploreResp, err := cfg.pokeapiClient.GetLocationArea(args[0])
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range exploreResp.PokemonEncounters {
		fmt.Println(" - ", encounter.Pokemon.Name)
	}
	return nil
}
