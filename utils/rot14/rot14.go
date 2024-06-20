package main

import "os"

func main(){
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	n := os.Args[1]
	result := ""

	for _, v := range n {
		if v >= 'a' && v <= 'z' {
			result += string((v - 'a' + 14)%26 + 'a')


		} else if v >= 'A' && v <= 'Z' {
			result += string((v - 'A' + 14)%26 + 'A')
	} else {
		result += string(v)
	}
} 
os.Stdout.WriteString(result + "\n")
}