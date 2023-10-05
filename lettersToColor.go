package main

import (
	"os"
	"strconv"
	"strings"
)

// convert the strings into ints that we can use to colour letters
func lettersToColor() []int {
	var numbersToColor []int
	var letterNumbers string
	// chooses the correct arg. depending upon the number of args and whether or not a font is chosen
	if len(os.Args) == 5 {
		letterNumbers = os.Args[2]
	}
	if len(os.Args) == 4 && os.Args[len(os.Args)-1] != "standard" && os.Args[len(os.Args)-1] != "shadow" && os.Args[len(os.Args)-1] != "thinkertoy" {
		letterNumbers = os.Args[2]
	}
	// if arg has a "-" then choose numbers between that range including those numbers
	if strings.Contains(letterNumbers, "-") {
		stringNumbers := strings.Split(letterNumbers, "-")
		firstInt, _ := strconv.Atoi(stringNumbers[0])
		secondNumber, _ := strconv.Atoi(stringNumbers[1])
		for i := firstInt; i <= secondNumber; i++ {
			numbersToColor = append(numbersToColor, i)
		}
		// if arg has a " " then split by " " and return every number
	} else if strings.Contains(letterNumbers, " ") {
		stringNumbers := strings.Split(letterNumbers, " ")
		for i := 0; i < len(stringNumbers); i++ {
			intNumber, _ := strconv.Atoi(stringNumbers[i])
			numbersToColor = append(numbersToColor, intNumber)
		}

		// arg probably only has 1 number
	} else if len(letterNumbers) == 1 {
		intNumber, _ := strconv.Atoi(letterNumbers)
		numbersToColor = append(numbersToColor, intNumber)

	} else {
		// if all else fails then say that numbersToColor is 10000000, will use this in the ascii function to just colour every letter
		numbersToColor = []int{1000000}
	}
	return numbersToColor
}
