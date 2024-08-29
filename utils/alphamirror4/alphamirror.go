package main

import (
	"os"

	"github.com/01-edu/z01"
)
func main(){
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	n := os.Args[1]

	for _, char := range n {
		if char >= 'a'  && char <= 'z' {
			z01.PrintRune('z' - (char - 'a'))
		} else if char >= 'A' && char <= 'Z' {
			z01.PrintRune('Z' - (char - 'A'))
		} else {
			z01.PrintRune(char)
		}
	} 
	z01.PrintRune('\n')
	
}
