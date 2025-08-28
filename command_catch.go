package main

import (
	"fmt"
	"math/rand"
)

func catch(BaseExperience int) bool {
	// Simple catch formula: the higher the base experience, the harder to catch
	// Max base experience is 306 (Mewtwo), min is around 39 (Caterpie)
	catchThreshold := 30000 / BaseExperience

	randomNumber := rand.Intn(1000)

	return randomNumber < catchThreshold
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please provide a pokemon name to catch. ex: catch clefairy")
	}

	name := args[0]
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResp.Name)
	if catch(pokemonResp.BaseExperience) {
		fmt.Printf("%s was caught!\n", pokemonResp.Name)

		// Add to pokedex
		cfg.pokedex[pokemonResp.Name] = pokemonResp

		return nil
	}
	fmt.Printf("%s escaped!\n", pokemonResp.Name)

	return nil
}
