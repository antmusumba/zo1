package main

import (
	"os"

	"github.com/01-edu/z01"
)
func main() {
	if len(os.Args) != 2 {
		return
	}
 args:= os.Args[1]
runes := []rune(args)
for i:=len(runes)-1 ; i < len(runes) ; i++ {
if runes[i] != ' ' {
	w := i 
	for i < len(runes) && runes[i] != ' ' {
		i++
	}
	wrd := runes[w:i]
	printstr(string(wrd))
	break
	
}
}
z01.PrintRune('\n')
}

func printstr(s string) {
	for _, val := range s {
		z01.PrintRune(val)
	}
}
