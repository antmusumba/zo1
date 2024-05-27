package utils

import "github.com/01-edu/z01"


func AbCd(){
	for x := 'z' ; x >= 'a' ; x-- {
		if x % 2 == 0 {
			z01.PrintRune(x)
		} else {
			z01.PrintRune(x-32)
		}
	

}
}