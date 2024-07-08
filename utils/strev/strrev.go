package main

import (
	"fmt"
	"os"
)


func StrRev(s string) string{
	var y string
	for _, i := range s {
		y = string(i) + y
	}
	return y
}
func main(){
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	str := StrRev(os.Args[1])
	fmt.Println(str)
}