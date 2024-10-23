package main

import (
    "fmt"

    
)

func main() {
    fmt.Print(FirstWord("hello there"))
    fmt.Print(FirstWord(""))
    fmt.Print(FirstWord("hello   .........  bye"))
}
// func FirstWord(s string) string {
	
//     result := spplit(s)
// 	varr := ""
// 	for i := 0 ; i < len(result) ; i++    {
// 		varr = result[0]

// 	}
// 	return varr + "\n"
// }
// func spplit(s string) []string {
// 	slc := []string{}
// 	result := ""
// 	for i := 0 ; i < len(s) ; i++ {
// 		if s[i] !=  ' ' {
// 			result += string(s[i])
// 		} else if s[i] == ' ' && result != "" {
// 			slc = append(slc, result)
// 			result = ""
// 		}
// 	}
// 	if result != "" {
// 		slc = append(slc, result)
// 	}
// 	return slc
// }
func FirstWord(s string) string {
    first := ""
	for i := 0 ; i < len(s) ; i++ {
		if s[i] == ' ' {
			break
		}
		first += string(s[i])
	}
	return first + "\n"
}
