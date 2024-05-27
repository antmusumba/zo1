package utils

import "github.com/01-edu/z01"

//import "github.com/01-edu/z01"

func Itoa(num int){
	result := ""
	for num != 0{
		mod := num % 10
		startrune := '0'
		for i := 0; i < mod; i++{
			startrune++
		}
		result =  string(startrune) + result
		num = num / 10
	}
	
	for _, val := range result {
	z01.PrintRune(val)
	}
	z01.PrintRune('\n')
	
}