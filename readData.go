package main

import (
	"log"
	"os"
)

// reads the data from whichever file is specified by the banner then returns it as a string
func readData(banner string) string {
	data, err := os.ReadFile(banner)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
