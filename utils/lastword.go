package utils

import "github.com/01-edu/z01"

func Lastword(s string) {
	// "antony musumba"
	var res []rune
	srune := []rune(s)
	for i := len(srune) - 1; i >= 0; i-- {
		if Alphan(srune[i]) {
			res = append(res, srune[i])
		}
		if srune[i] == ' ' && len(res) > 0 {
			break
		}
	}
	for i := len(res)-1; i >= 0; i-- {
		z01.PrintRune(res[i])
	}
	if len(res) == 0 {
		return
	}
	z01.PrintRune('\n')
}

func Alphan(s rune) bool {
	if s > ' ' && s <= '~' {
		return true
	}
	return false
}
