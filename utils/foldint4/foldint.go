package main

import (
	"fmt"
)

// Mul multiplies two integers
func Mul(a, b int) int {
	return a * b
}

// Sub subtracts the second integer from the first
func Sub(a, b int) int {
	return a - b
}

// Add adds two integers
func Add(a, b int) int {
	return a + b
}


func FoldInt(f func(int, int) int, a []int, n int) {
	for _, val := range a {
		n = f(n, val)
	}
	fmt.Println(n)
}

func main() {
	table := []int{1, 2, 3}
	ac := 93

	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}
