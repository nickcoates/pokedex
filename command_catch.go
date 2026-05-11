package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", args[0])

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return errors.New("this is not a real pokemon")
	}

	chance := rand.IntN(pokemonResp.Experience)
	threshold := 40

	if chance <= threshold {
		fmt.Printf("%v was caught!\n", args[0])
		cfg.caughtPokemon[pokemonResp.Name] = pokemonResp
	} else {
		fmt.Printf("%v escaped!\n", args[0])
	}

	return nil
}
