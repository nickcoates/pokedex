package main

import "strings"

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	splitstrings := strings.Fields(lowered)
	return splitstrings
}
