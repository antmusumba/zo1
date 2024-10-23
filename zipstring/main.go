package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dogggg"))
	fmt.Println(ZipString("Helloo Therre!"))
}
func ZipString(s string) string {
	count := 1 
	res := ""
	for i := 0 ; i < len(s) ; i++ {
		if i + 1 < len(s) {
			if s[i] == s[i+1] {
				count++
			} else {
				res += strconv.Itoa(count) + string(s[i])
				count = 1

			}
		}
	}
	return res + strconv.Itoa(count) + string(s[len(s)-1])

}
