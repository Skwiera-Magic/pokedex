package main

import "fmt"

func commandHelp() error {
	fmt.Print("\nWelcome to the Pokedex!")
	fmt.Print("\nUsage:\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Print("\n")
	return nil
}