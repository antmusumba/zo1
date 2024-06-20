package main

import (
	"os"
	"strconv"
	//"zo1/utils"

	"github.com/01-edu/z01"
)
func main () {
	args, _ := strconv.Atoi(os.Args[1])
	str := ""

	for i := 1 ; i <= 9 ; i++ {
		x := i * args
		str += strconv.Itoa(i) + " x " + strconv.Itoa(args) + " = " + strconv.Itoa(x) + "\n"
	}
	for _,y := range str {
		z01.PrintRune(y)
    }
	z01.PrintRune('\n')

	
}



