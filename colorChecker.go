package main

import (
	"os"
	"strings"
)

// checks to see if a colour is specified in the first arg, if it is then returns the colour code
func colorChecker() string {
	color := "\u001b[37m"
	if strings.Contains(os.Args[1], "black") {
		// color := os.Args[1]
		color := "\u001b[30m"
		return color

	} else {
		if strings.Contains(os.Args[1], "red") {
			// color := os.Args[1]
			color := "\u001b[31m"
			return color

		} else {
			if strings.Contains(os.Args[1], "green") {
				// color := os.Args[1]
				color := "\u001b[32m"
				return color

			} else {
				if strings.Contains(os.Args[1], "yellow") {
					// color := os.Args[1]
					color := "\u001b[33m"
					return color

				} else {
					if strings.Contains(os.Args[1], "blue") {
						// color := os.Args[1]
						color := "\u001b[34m"
						return color

					} else {
						if strings.Contains(os.Args[1], "magenta") {
							// color := os.Args[1]
							color := "\u001b[35m"
							return color
						} else {
							if strings.Contains(os.Args[1], "cyan") {
								// color := os.Args[1]
								color := "\u001b[36m"
								return color
							} else {
								if strings.Contains(os.Args[1], "orange") {
									color := "\033[38:5:208m"

									return color

								}
							}
						}
					}
				}
			}
		}
	}
	return color
}
