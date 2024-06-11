package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args[1:]) != 2 {
		fmt.Println("provide enough arguments")
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]

	i := 0
	j := 0

	for i < len(str1) && j < len((str2)) {
		if str1[i] == str2[j] {
			i++
		}
		j++

	}
	if i == len(str1) {
		fmt.Println(str1)
	} else {
		return
	}







}