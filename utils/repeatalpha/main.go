package main

import (
	"fmt"
	
)

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
func RepeatAlpha(s string) string {
	res := ""
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			index := 0
			if char >= 'A' && char <= 'Z' {
				index = int(char) - 'A' + 1
			} else {
				index = int(char) - 'a' + 1
			}
			for i := 0; i < index ; i++ {
				res += string(char)
			}
		} else {
			res += string(char)
		}
	}
	return res
}
