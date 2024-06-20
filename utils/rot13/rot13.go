package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	n := os.Args[1]
	newstr := ""

	for _, v := range n {
	if v >= 'a' && v <= 'z' {
		newstr += string((v - 'a' + 13)%26 + 'a')
	} else if v >= 'A' && v <= 'Z' {
		newstr += string((v - 'A' + 13)%26 + 'A')
	} else {
		newstr +=string(v)
	}
		
	}
	//os.Stdout.WriteString(newstr + "\n")
	fmt.Println(newstr)
}
