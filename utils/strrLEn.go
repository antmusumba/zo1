package utils

import "github.com/01-edu/z01"

func StrLen(s string) {
	// count := len(s) + '0'
	// z01.PrintRune(rune(count))
	// z01.PrintRune('\n')
	coun := '0'
	for range s {
		coun++
	} 
	z01.PrintRune(coun)
	z01.PrintRune('\n')
}
