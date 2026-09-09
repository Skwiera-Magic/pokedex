package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Print("\nWelcome to the Pokedex!")
	fmt.Print("\nUsage:\n")
	for _, cmd := range cfg.commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Print("\n")
	return nil
}