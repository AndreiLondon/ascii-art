package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	text := "hello"
	want := `
 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                              `
	result := convertText(text)
	if want != result {
		t.Fatalf(`convertText(text) fail.
		Want:  %v 
		Found: %v`, want, result)
	}
}

func convertText(text string) string {
	path := "standard.txt"
	output := readFile(path)
	myMap := parseTemplate(output)
	return printMessageToString(text, myMap)
}

func printMessageToString(s string, myMap map[rune][][]rune) string {
	result := " " + "\n"

	split := strings.Split(s, "\\n")

	// We need this check when we pass \n twice in a row, so we don't need to print an empty string
	for _, v := range split {
		if v == "" {
			fmt.Println()
			continue
		}
		// 1st loop is for counting the lines
		for i := 0; i < 8; i++ {
			// 2nd loop is for grabbing the each char from the string
			for _, ch := range v {
				// do whatever you want to handle errors - this means this wasn't a string
				if data, ok := myMap[ch]; ok {
					rowData := data[i]
					for _, value := range rowData {
						// fmt.Print(string(value))

						result = result + string(value)
					}
				}
			}
			result = result + string('\n')
		}

	}
	return result
}
