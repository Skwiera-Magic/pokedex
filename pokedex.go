package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
)

func pokedex(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanned := scanner.Scan()
		if scanned == false {
			scanner.Err()
			break
		}
		texted := scanner.Text()
		cleaned := cleanInput(texted)
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}
		
		command, exists := cfg.commands[cleaned[0]]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Print("Unknown command, type help\n")
		}
	}
}

type config struct {
	commands 			map[string]cliCommand
	pokeapiClient 		pokeapi.Client
	nextLocationsURL 	*string
	prevLocationsURL 	*string
}

type cliCommand struct {
	name			string
	description		string
	callback		func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: 			"help",
			description: 	"Displays help menu",
			callback: 		commandHelp,
		},
		"map": {
			name: 			"map",
			description: 	"Next locations page",
			callback: 		commandMapf,
		},
		"mapb": {
			name: 			"mapb",
			description: 	"Previous locations page",
			callback: 		commandMapb,
		},
		"explore": {
			name: 			"explore <location_name>",
			description: 	"Check what pokemon are in the location",
			callback: 		commandExplore,
		},
		"exit": {
			name: 			"exit",
			description: 	"Exit the Pokedex",
			callback: 		commandExit,
		},
	}
}