package utils

import "github.com/01-edu/z01"

func PrintHello(){

s := "helloWorld"
for _, val := range s {
	z01.PrintRune(val)

}
z01.PrintRune('\n')
}