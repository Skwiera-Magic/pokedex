package main

func main() {
	cfg := &config{
		commands: getCommands(),
	}
	pokedex(cfg)
}