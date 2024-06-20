package main

import (
	"fmt"
	
	"zo1/utils"
)

func FindPrevPrime(nb int) int {
   for i:= nb; i >= 0; i--{
	if utils.Isprime(i) {
		return i
	}
   }
   return 0
}

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}


