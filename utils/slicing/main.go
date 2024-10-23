// package main

// import "fmt"
// func main(){
// 	str :=  "Hello, 世界q"
// 	ru := []rune(str)

// 	// i := 7
// 	// end := 11
// 	// result := []rune{}
// 	result:= ru[7:11]
// 	//result = append(result, []rune(str[i:end])...)
// 	fmt.Println(string(result))
// }
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
	if num <= 0 {
		return arg
	}
	save := true
	result := []rune{}
	for i := 0 ; i < len(arg) ; i+=num {
		if save {
			end := i + num 
			if end > len(arg) {
				end = len(arg)
			}
			result = append(result, []rune(arg[i:end])...)
		}
		save = !save
	}
	return string(result)
}