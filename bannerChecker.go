package main

import "os"

// Checks to see if the final argument is one of the three fonts
func bannerChecker() string {
	argsLen := len(os.Args) - 1
	if os.Args[argsLen] == "shadow" || os.Args[argsLen] == "standard" || os.Args[argsLen] == "thinkertoy" {
		// if so return that font
		banner := os.Args[argsLen]
		return banner
	} else {
		// otherwise return "standard", acting as a default
		banner := "standard"
		return banner
	}
}
