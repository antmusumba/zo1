package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) != 2 {
		os.Exit(0)
	}
	a := os.Args[1]
	r := []rune(a)
	HasItBreak := false
	var lastword []rune
	for i := len(r) - 1; i >= 0; i-- {

		v := r[i]
		// "  lorem,ipsum  "
		// lorem     ipsum      "
		if v == ' ' && HasItBreak {
			break
		} else if v == ' ' && !HasItBreak {
			continue
		} else {
			lastword = append(lastword, v)
			HasItBreak = true

		}

	}
	for i := len(lastword) - 1; i >= 0; i-- {
		z01.PrintRune(lastword[i])
	}
	if len(lastword) != 0 {
		z01.PrintRune('\n')
	}
}
