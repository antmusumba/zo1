package main

import (
	"fmt"
	

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

func FindPrevPrime(nb int) int {
   for i:= nb; i >= 0; i--{
	if Isprime(i) {
		return i
	}
   }
   return 0
}

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}


