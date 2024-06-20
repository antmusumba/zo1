package main

import (
	"zo1/utils"

	"github.com/01-edu/z01"
)

func ReduceInt(a []int, f func(int, int) int) {
	n := a[0]
	for i := 1; i <= len(a)-1; i++ {
		n = f(n, a[i])
	}
	
	v := utils.Itoa(n)
	for _, ch := range v {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}
