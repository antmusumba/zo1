package main

import (
	"fmt"
	"os"
	"zo1/utils"
)

func Isprime(num int) bool {
	if num < 2 {
		return false
	}
	if num == 2 || num == 3 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	
	 for i := 2; i*i <= num; i += 2 {
	 	if num%i == 0 {
	 		return false
	 	}
	 }
	return true
}
func main(){
	if len(os.Args) !=  2 {
		os.Exit(1)
	}
	str := Isprime(utils.BasicAtoi(os.Args[1]))
	fmt.Printf("Number %s is %v : prime number\n", os.Args[1], str)
}
