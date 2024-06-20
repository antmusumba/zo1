package main

import (
	//"fmt"
	"os"
	//"zo1/utils"

	"github.com/01-edu/z01"
)

func main(){
 if len(os.Args) != 2{
	return
 }

	args  := os.Args[1]
	for _, char := range args {
		if char >= 'a'  && char <= 'z' {
			 z01.PrintRune('z' - (char-'a'))
	}else if char >= 'A' && char <= 'Z' {
		z01.PrintRune('Z'- (char- 'A'))
	}else{
		z01.PrintRune(char)
	}
	
}
z01.PrintRune('\n')
}