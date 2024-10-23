package main

import (
	"fmt"
	
)

func main() {
	fmt.Println(FishAndChips(4))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
}
func FishAndChips(n int) string {
	result := ""

	if n % 2 == 0 {
		result = "fish"
	}
	if n % 3 == 0 {
		result = "chips"
	}
	if n % 2 == 0 && n % 3 == 0 {
		result = "fish and chips"
	}
	if n < 0 {
		result = "error: number is negative"
	}
	if n % 2 != 0 && n % 3 != 0 {
		result = "error: non divisible"
	}
	return result

}
