package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s1 := os.Args[1]
	s2 := os.Args[2]
	result := ""
	seen := make(map[rune]bool)
	for _, v := range s1 {
		if !seen[v] {
			for _, ch := range s2 {
				if v == ch {
					result += string(v)
					seen[v] = true
					break

				}
			}
		}
	}
	for _, ch := range result {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
