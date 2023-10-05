package main

import (
	"fmt"
	"strings"
)

// intakes a color, input, banner (font) and 'letter numbers to be coloured'.
// if color, banner or letter numbers are not specified in the arguments then programme will use default values for each
func asciiColorFS(color string, input string, banner string, lettersToColor []int) {
	// Text colour = color
	// Text content = input
	// Text banner/font = banner

	// the chosen banner, split by '\n'
	splitStr := strings.Split(string(readData(banner+".txt")), "\n")
	// input, split by new line; allows \n to print input on multiple lines
	splitInput := strings.Split(string(input), "\\n")
	// for each chunk of the input (split by new line)
	for a := 0; a < len(splitInput); a++ {
		// divide it by individual rune (letter)
		runeInput := []rune(splitInput[a])
		// if the slice of arg contains nothing then print a new line
		if len(splitInput[a]) == 0 {
			(print("\n"))
		} else {
			// otherwise print the first line of each letter in the slice of arg, then 2nd ... 8th
			// i = the line of the ascii character
			for i := 1; i <= 8; i++ {
				// j = the runes of the input, cycled through one at a time
				for j := 0; j <= len(runeInput)-1; j++ {
					letterInput := runeInput[j]
					// rune to line number
					lineNumber := (int(letterInput)-32)*9 + i
					// prints from line number to the next 8 lines
					// ColorReset after desired strings to turn text back to white
					ColorReset := "\u001b[0m"
					// if j == any of the letter numbers to colour then it'll colour it.  If it equals 1000000 (colour specified, but no letters specified) then it'll colour every letter
					if jIsColoured(j, lettersToColor) || lettersToColor[0] == 1000000 {
						fmt.Print((color), (splitStr[lineNumber]), string(ColorReset))
					} else {
						// otherwise it prints without colour
						fmt.Print(splitStr[lineNumber])
					}

				}
				fmt.Println()
			}
		}
	}
}

// checks to see if the letter position that is being printed should be printed in colour
func jIsColoured(j int, lettersToColor []int) bool {
	for i := 0; i < len(lettersToColor); i++ {
		if j == lettersToColor[i]-1 {
			return true
		}
	}
	return false
}
