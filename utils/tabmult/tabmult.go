package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)
func main () {
	args , _ := strconv.Atoi(os.Args[1])
	str := ""
	for i := 1 ; i < 9 ; i++ {
		x := i * 10
		str += strconv.Itoa(i) + " + " + strconv.Itoa(args) + " = " + strconv.Itoa(x) + "\n"
	}
	for _,y := range str {
		z01.PrintRune(y)
    }
	z01.PrintRune('\n')

	
}




