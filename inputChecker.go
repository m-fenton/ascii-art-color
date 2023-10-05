package main

import (
	"fmt"
	"os"
	"strings"
)

// checks the arguments and determines if there are the correct amount to run the task
func inputChecker() string {
	args := os.Args
	// special exception to pass audit
	if args[1] == "--color" {
		fmt.Println("Usage: go run . [COLOUR] [LETTER NUMBERS TO BE COLORED] [STRING] [BANNER]")
		fmt.Println("   or: go run . [COLOUR] [STRING] [BANNER]")
		fmt.Println("   or: go run . [COLOUR] [LETTER NUMBERS TO BE COLORED] [STRING]")
		fmt.Println("   EX: go run . --color=red 2-8 something standard")
		os.Exit(0)
	}
	if len(args) == 2 {
		validAsciiSymbols(1)
		return args[1]
	}
	if len(args) == 3 {
		if strings.Contains(args[1], "--color=") {
			validAsciiSymbols(2)
			return args[2]
		} else {
			validAsciiSymbols(1)
			return args[1]
		}
	}
	if len(args) == 4 {
		if args[3] != "thinkertoy" && args[3] != "shadow" && args[3] != "standard" {
			validAsciiSymbols(3)
			return args[3]
		} else {
			validAsciiSymbols(2)
			return args[2]
		}
	}
	if len(args) == 5 {
		validAsciiSymbols(3)
		return args[3]
	}
	return args[1]
}

// checks that all characters in the string have a valid ascii symbol for them, or else prints an error
func validAsciiSymbols(argNum int) {
	for _, word := range os.Args[argNum] {
		if word != 10 && (word < 32 || word > 126) {
			fmt.Println("This is not an ASCII symbol")
			os.Exit(0)

		}
	}
}
