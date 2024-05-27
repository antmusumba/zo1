package utils

import (
	//"fmt"

	//"github.com/01-edu/z01"
)


func BasicAtoi(s string) int {
	// x := 0

	// for _, val := range s {
	// 	y := 0
	// 	for i := '1' ; i <= val ; i++ {

	// 		y++
	// 	}
	// 	x = x*10 + y
		
	// }
	// return x

	y := 0
	for _,val := range s {
		y =y*10 + int(val-'0')
	}
	return y

}