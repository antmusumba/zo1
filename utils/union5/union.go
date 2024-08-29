package main

import (
	"fmt"
	"os"
)

// uniqueCharacters returns a string of unique characters from both input strings
func uniqueCharacters(s1, s2 string) string {
	available := make(map[rune]bool)
	result := ""
	for _ , v := range s1 + s2 {
		if !available[v] {
			available[v] = true
			result += string(v)

		}
	}
	return result

}

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	unique := uniqueCharacters(s1, s2)
	fmt.Println(unique)
}
