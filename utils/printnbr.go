package utils

import "github.com/01-edu/z01"

func PrintNbr(s int) {
	n := s

	//z01.PrintRune(rune(n))
	var r []rune
for n > 0 {
	m := n%10
	n /=10
	r = append(r, rune(m+'0'))
}
for i := len(r)-1; i >= 0; i-- {
	z01.PrintRune(r[i])
}
	z01.PrintRune('\n')
}
