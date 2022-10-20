package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	//"strings"
	//"io/ioutil"

	"os"
)

func main() {

	// wordArg := os.Args
	remove := strings.ReplaceAll(os.Args[1], "\\n", "\n")
	//os.Args[1] argument or maybe some others instance is the given string
	//"\n" is the string which we want to replace, and
	//"\n" is string which will replace the old string.
	fmt.Println(remove)
	// wordRune := []rune(os.Args[1])
	// fmt.Println(wordRune)
	wordRune := []rune(remove)
	fmt.Println(wordRune)

	f, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 8; i++ {
		//if i == \ && [i+1] == n {
		//    fmt.Println()
		//}
		for j := 0; j < len(wordRune)-1; j++ {
			if lines[int(wordRune[j])*9-287+i] == "        " {
				fmt.Printf("        ")
			} else {
				fmt.Printf(lines[int(wordRune[j])*9-287+i])
			}
		}
		fmt.Print("\n")
	}

}
