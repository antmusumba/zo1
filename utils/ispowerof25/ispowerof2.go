package main

import "os"

func isPowerOfTwo(n int) bool {
	return n > 0 && (n & (n -1)) == 0
}

func atoi(str string) int {
	num := 0
	neg := false 
	if str[0] == '-'{
		neg = true
		str = str[1:]

	}
	for _, v := range str {
		if v < '0' || v > '9' {
			os.Exit(1)
		}
		num = num*10 + int(v -'0')
	}
	if neg {
		num = -num
	}
	return num
}

 func main(){
	if len(os.Args) != 2 {
		return
	}
	num := atoi(os.Args[1])
	if isPowerOfTwo(num) {
		os.Stdout.WriteString("true$" + "\n")
	} else {
		os.Stdout.WriteString("false$" + "\n")
	}


 }