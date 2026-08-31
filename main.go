package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
		fmt.Printf("Your command was: %s\n", cleaned[0])
	}
}