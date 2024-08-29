package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if the number of arguments is 3
	if len(os.Args) != 4 {
		return
	}

	first := os.Args[1]
	second := os.Args[2]
	third := os.Args[3]

	// Ensure the second and third arguments are single characters
	if len(second) != 1 || len(third) != 1 {
		return
	}

	// Check if the oldChar is in the inputString
	found := false
	for _, v := range first {
		if string(v) == second {
			found = true
			break

		}
	}

	if !found {
		fmt.Println(first)
		return
	}

	// Replace oldChar with newChar in the inputString
	outputString := ""
	for _, v := range first {
		if string(v) == second {
			outputString += third

		} else {
			outputString += string(v)
		}
	}

	fmt.Println(outputString)
}
