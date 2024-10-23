package main

import "fmt"

func main() {
	fmt.Println(ThirdTimeIsACharm("123456789"))
	fmt.Println(ThirdTimeIsACharm(""))
	fmt.Println(ThirdTimeIsACharm("a b c d e f"))
	fmt.Println(ThirdTimeIsACharm("12"))
}

func ThirdTimeIsACharm(str string) string {
	res := ""
	for i := 2; i < len(str); i += 3 {
		res += string(str[i])
	}
	return res
}
