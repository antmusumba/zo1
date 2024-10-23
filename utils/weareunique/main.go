package main

import (
	"fmt"
)

func main() {
	str1 := "foo"
	str2 := "boo"
	fmt.Println(weAreUniq(str1, str2))
}

func weAreUniq(a, b string) (final int) {
	uniQ := make(map[rune]int)
	for _ ,ch := range a {
		uniQ[ch]++
	}
	for _, ch := range b {
		uniQ[ch]++
	}
	for _ , count := range uniQ {
		if count == 1 {
			final++
		}
	}
	return final

}
