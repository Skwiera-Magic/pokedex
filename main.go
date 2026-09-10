package main

import (
	"pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, time.Minute * 5)
	cfg := &config{
		commands: 		getCommands(),
		caughtPokemon:	map[string]pokeapi.Pokemon{},
		pokeapiClient: 	pokeClient,
	}
	pokedex(cfg)
}