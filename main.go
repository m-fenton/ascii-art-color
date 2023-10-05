package main

import (
	"fmt"
	"os"
)

// working on letterstocolor

// Reads the text file specified by the banner

func main() {
	if len(os.Args) == 1 ||
		// len(os.Args) == 2 ||
		len(os.Args) > 5 {
		fmt.Println("Usage: go run . [COLOUR] [LETTER NUMBERS TO BE COLORED] [STRING] [BANNER]")
		fmt.Println("   or: go run . [COLOUR] [STRING] [BANNER]")
		fmt.Println("   or: go run . [COLOUR] [LETTER NUMBERS TO BE COLORED] [STRING]")
		fmt.Println("   EX: go run . --color=red 2-8 something standard")
		os.Exit(0)

	} else {

		color := colorChecker()
		input := inputChecker()
		banner := bannerChecker()
		letters := lettersToColor()

		asciiColorFS(color, input, banner, letters)
		//}

	}
}
