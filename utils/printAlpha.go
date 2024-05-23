package utils

import "github.com/01-edu/z01"


func PrintAlphabets(){
	for i:= 'A' ; i <= 'Z' ;i++ {
		z01.PrintRune(i)
	}
}