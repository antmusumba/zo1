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

	fst := os.Args[1]
	scd := os.Args[2]
	trd := os.Args[3]

	// Ensure the second and third arguments are single characters
	if len(scd) != 1 || len(trd) != 1 {
		return
	}

	// Check if the oldChar is in the inputString
	found := false
	for _, v := range fst {
		if string(v) == scd {
			found = true
			break

		}
	}

	if !found {
		fmt.Println(fst)
		return
	}

	// Replace oldChar with newChar in the inputString
	outputString := ""
	for _, v := range fst {
		if string(v) == scd {
			outputString += trd

		} else {
			outputString += string(v)
		}
	}

	fmt.Println(outputString)
}
