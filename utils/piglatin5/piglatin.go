package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	vowel := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

	str := os.Args[1]
	hasVowel := false

	for i, v := range str {
		for _, char := range vowel {
			if v == char {
				str = str[i:] + str[:i]
				hasVowel = true
				break

			}
		}
	}
	if !hasVowel {
		fmt.Println("No vowels")
	}

	for _, char := range vowel {
		if rune(str[0]) == rune(char) {
			str += "ay"
			fmt.Println(str)

		}
	}
}
