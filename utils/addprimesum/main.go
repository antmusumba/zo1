package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("not enough arguments")
	}
	res := 0
	num := ato(os.Args[1])
	for i := 2 ; i <= num ; i++ {
		if Iprime(i) {
			res += i
		}
	}
	fmt.Println(res)
}
func ato(s string) int {
	res := 0
	for i := 0 ; i < len(s) ; i++ {
		res = res*10 + int((s[i]) - '0')
	}
	return res
}

func Iprime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 2 ; i*i <= n ; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
