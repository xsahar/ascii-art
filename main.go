package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide words to print.")
		return
	}

	input := os.Args[1]
	if input == "\\n" {
		fmt.Println()
		return
	} else if input == "" {
		return
	}

	data, err := os.ReadFile("standard.txt")
	if err != nil {
		return
	}

	lines := strings.Split(string(data), "\n")
	words := strings.Split(input, "\n")

	for _, word := range words {
		if word == "" {
			fmt.Println()
			continue
		}

		printWord(word, lines)
	}
}

func printWord(word string, lines []string) {
	for j := 1; j < 9; j++ {
		str := ""

		for _, letter := range word {
			val := int(letter)
			line := (val - 32) * 9
        	if line+j >= len(lines) {
				return
			}

			str += lines[line+j]
		}

		fmt.Println(str)
	}
}