package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	res := []string{}
	for _, ch := range arr {
		if ch == 0 {
			res = append(res, "00")
		} else {
			res = append(res, Hex(int(ch)))
		}
	}

	for i, str := range res {
		for _, ch := range str {
			z01.PrintRune(ch)
		}
		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		} else if i != len(res)-1 {
			z01.PrintRune(' ')
		}

	}
	z01.PrintRune('\n')
	for _, ch := range arr {
		if ch >= ' ' && ch <= '~' {
			z01.PrintRune(rune(ch))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func Hex(num int) string {
	result := ""
	hexvalues := "0123456789abcdef"
	for num > 0 {
		mod := num % 16
		result = string(hexvalues[mod]) + result
		num /= 16
	}
	return result
}
