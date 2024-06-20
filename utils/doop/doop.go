package main

import (
	"os"
	
	"zo1/utils"

)

func main() {
	if len(os.Args) != 4 {
		os.Stdout.WriteString("not enough arguments" + "\n")
		os.Exit(0)
	}

	// Parse first value
	firstValue := os.Args[1]
	if len(firstValue) > 11 {
		return
	}
	f := utils.BasicAtoi(firstValue)
	
	f32 := int32(f)

	// Parse third value
	thirdValue := os.Args[3]
	if len(thirdValue) > 11 {
		return
	}
	t := utils.BasicAtoi(thirdValue) 
	
	
	t32 := int32(t)

	// Parse operator
	operator := os.Args[2]
	var answer int32

	switch operator {
	case "+":
		answer = f32 + t32
	case "-":
		answer = f32 - t32
	case "/":
		if t32 == 0 {
			os.Stdout.WriteString("No division by 0" + "\n")
			os.Exit(0)
		}
		answer = f32 / t32
	case "*":
		answer = f32 * t32
	case "%":
		if t32 == 0 {
			os.Stdout.WriteString("No modulo by 0" + "\n")
			os.Exit(0)
		}
		answer = f32 % t32
	default:
		os.Stdout.WriteString("invalid operator" + "\n")
		os.Exit(0)
		
	}


	ans := utils.Itoa(int(answer))
	os.Stdout.WriteString(ans + "\n")
	
}
