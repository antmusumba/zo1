package utils

import (
	"fmt"
	"os"
)

//"fmt"

//"github.com/01-edu/z01"

func BasicAtoi(s string) int {

	if s == "0" {
		return 0
	}
	neg := false
	// -4894
	if s[0] == '-' {
		neg = true
		s = s[1:]
	}

	y := 0
	for _, val := range s {
		if val < '0' || val > '9' {
			fmt.Println("cannot convert")
			os.Exit(1)
		}
		y = y*10 + int(val-'0')
	}
	if neg {
		y = -y
	}
	return y
}

