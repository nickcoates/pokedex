package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type config struct {
	Next     *string
	Previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type locationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type locationAreaResp struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []locationArea `json:"results"`
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := config{}
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned_text := cleanInput((text))[0]
		commands := getCommands()
		command, exists := commands[cleaned_text]
		if exists {
			err := command.callback(&cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, value := range getCommands() {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	var url string
	if cfg.Next != nil {
		url = *cfg.Next
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var data locationAreaResp
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	cfg.Next = data.Next
	cfg.Previous = data.Previous

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	var url string
	if cfg.Previous != nil {
		url = *cfg.Previous
	} else {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var data locationAreaResp
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	cfg.Next = data.Next
	cfg.Previous = data.Previous

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 areas in the Pokemon world",
			callback:    commandMapb,
		},
	}
}
