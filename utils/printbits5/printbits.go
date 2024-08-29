package main

import (
	"os"
	"strconv"
)

func main()  {
	args := os.Args[1]
	myargs:= atoi(args) 
	if myargs < 0 || myargs > 255 {
		os.Stdout.WriteString("00000000$")
		return
	}

	new := strconv.FormatUint(uint64(myargs),2)
	for len(new) < 8 {
		new = "0" + new
	}

	os.Stdout.WriteString(new + "\n")

}
func atoi(s string) int {
	if s == "0" {
		return 0
	}
	num := 0
	neg := false
	for _, v := range s {
		if v < '0' || v > '9' {
			os.Exit(0)
		}
		num = num * 10 + int(v - '0')
	}
	if neg {
		num = -num
	}
	return num

}