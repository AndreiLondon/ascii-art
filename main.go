package main

import (
	"fmt"
)

func main() {
	result := getArguments()
	// If it's an empty string it will exit the program
	if result == "" {
		return
	}
	if result == "\\n" {
		fmt.Println()
		return
	}

	path := "standard.txt"
	output := readFile(path)
	myMap := parseTemplate(output)
	printMessage(result, myMap)
}
