package main

import (
	"fmt"
	"os"
	"zo1/utils"
)
 var sample = []struct {
	value int
 	symbol string
 } {
 {1000, "M"},
 {900, "(M-C)"},
 {500, "D"},
 {400, "(D-C)"},
 {100, "C"},
 {90, "(C-X)"},
 {50, "L"},
 {40, "(L-X)"},
 {10, "X"},
 {9, "(X-I)"},
 {5, "V"},
 {4, "(V-I)"},
 {1, "I"},
	
 }

 func main() {
 	if len(os.Args) != 2 {
 		os.Exit(0)
 	}
 	num := utils.BasicAtoi(os.Args[1])
	
 	var1, var2 := Romaniser(num)
	fmt.Println(var1)
	fmt.Println(var2)

 }
 func Romaniser(num int) (string,string) {
	result := ""
	steps := ""
 	for _, v := range sample {
 		for num >= v.value {
			if len(result) > 0 {
				steps += "+"
			}
			if v.symbol[0] == '(' {
				steps += v.symbol
				result += string(v.symbol[1]) + string(v.symbol[3])

			} else {
				steps = v.symbol
				result += v.symbol
			}
			num -=v.value
			
		}
 	}
	return steps,result


 }