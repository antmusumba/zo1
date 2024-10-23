package main

import (
	"fmt"
)

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}
func CountAlpha(s string) int {
	count := 0
	for i := 0 ; i < len(s) ; i++     {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z'{
			count ++
		}
	}
	return count

}

// func splash( s string) []string {
// 	slc := []string{}
// 	result := ""
// 	for i := 0 ; i < len(s) ; i++ {
// 		if s[i] != ' ' {
// 			result += string(s[i])
// 		} 
// 	}
// }