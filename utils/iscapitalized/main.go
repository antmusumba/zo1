package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}


func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] >= 'a' && s[0] <= 'z' {
		return false
	}
	for i := 1 ; i < len(s) ; i++ {
		if s[i] == ' ' && i+1 < len(s) {
			if s[i+1] >= 'a' && s[i+1] <= 'z' {
				return false
			}
		}
	}
	return true

}

