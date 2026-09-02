package main

import (
	"bufio"
	"os"
	"fmt"
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
		
		command, exists := cfg.commands[cleaned[0]]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Print("Unknown command\n")
		}
	}
}

type config struct {
	commands map[string]cliCommand
}

type cliCommand struct {
	name			string
	description		string
	callback		func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: 			"help",
			description: 	"Displays help menu",
			callback: 		commandHelp,
		},
		"exit": {
			name: 			"exit",
			description: 	"Exit the Pokedex",
			callback: 		commandExit,
		},
	}
}