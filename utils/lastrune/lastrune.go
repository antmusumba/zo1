package main

import "github.com/01-edu/z01"

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}

func LastRune(s string) rune {
	r := []rune(s)
	var lastrune rune

	for i := 0; i <= len(r)-1; i++ {
		lastrune = r[len(r)-1]
	}
	return lastrune
}
