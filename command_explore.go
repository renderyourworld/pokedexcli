package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please provide a location to explore. ex: explore pastoria-city-area")
	}

	name := args[0]
	locationResp, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", args[0])

	if len(locationResp.PokemonEncounters) != 0 {
		fmt.Println("Found Pokemon:")
	}

	for _, loc := range locationResp.PokemonEncounters {
		fmt.Println("-", loc.Pokemon.Name)
	}

	return nil
}
