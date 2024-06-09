package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args[1:]) < 1 || len(os.Args[1:]) > 2 {
		fmt.Println("usage : go run .  string")
		return
	}
	a := os.Args[1]
	r := []rune(a)
	var isBreak bool = false
	var lastword []rune
	for i := len(r) - 1; i >= 0; i-- {
		y := r[i]
		if y == ' ' && isBreak {
			break
		} else if y == ' ' && !isBreak {
			continue
		} else {
			lastword = append(lastword, y)
			isBreak = true

		}
	}
	for i := len(lastword) - 1; i >= 0; i-- {
		z01.PrintRune(lastword[i])
	}
	if len(lastword) != 0 {
		z01.PrintRune('\n')
	}
}
