package main

import (
	"fmt"
	"log"
	"os"
)

func getArguments() string {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("invalid number of arguments")
		//os.Exit closes the application and with 2 it means that
		os.Exit(2)
	}
	return args[1]
}

func readFile(s string) string {

	data, err := os.ReadFile(s)
	if err != nil {
		log.Fatal("invalid file")
	}
	return string(data)

}
