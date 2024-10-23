package main

import "fmt"

func main() {
	fmt.Println(SaveAndMiss("123456789", 3))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))
}
func SaveAndMiss(arg string, num int) string {
	if num == 0 {
		return arg
	}
	var saved []rune
	for i, char := range arg {
		if i % (2*)
	}
}