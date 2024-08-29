package main

import (
	"os"

	"github.com/01-edu/z01"
)

func StrLen(s string) {
	count := 0
	for range s {
		count++
	}
	newS := Ito(count)
	for _, ch := range newS {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func Ito(num int) string {
	result := ""
	if num == 0 {
		return "0"
	}
	neg := false
	if num < 0 {
		neg = true
		num = -num

	}
	for num > 0 {
		result = string(rune(num%10+'0')) + result
		num /= 10
	}
	if neg {
		result = "-" + result
	}
	return result
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	str := os.Args[1]
	StrLen(str)
}
