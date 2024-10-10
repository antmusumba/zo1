package main

import "fmt"

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
func HashCode(dec string) string {
	hashed := ""
	strlen := len(dec)
	for i := 0 ; i < strlen ; i++ {
		hashValue := (int(dec[i]) + strlen) % 127
		if hashValue < 32 {
			hashValue += 33
		}
		hashed += string(rune(hashValue))
	}
	return hashed
}
