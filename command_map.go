package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
	locationAreaResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationUrl = locationAreaResp.Next
	cfg.prevLocationUrl = locationAreaResp.Previous

	for _, loc := range locationAreaResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationUrl == nil {
		return errors.New("you're on the first page")
	}

	locationAreaResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationUrl = locationAreaResp.Next
	cfg.prevLocationUrl = locationAreaResp.Previous

	for _, loc := range locationAreaResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
