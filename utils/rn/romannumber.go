package main

import (
	"fmt"
	"os"
	"zo1/utils"
)

// Roman numeral mappings
var romanStruct = []struct {
	value int
	symbol string
}{
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


// Convert to Roman numeral with detailed steps
func intToRoman(n int) (string, string) {
	result := ""
	steps := ""
	for _, numeral := range romanStruct {
		for n >= numeral.value {
			if len(result) > 0 {
				steps += "+"
			}
			if numeral.symbol[0] == '(' {
				steps += numeral.symbol
				result += string(numeral.symbol[1]) + string(numeral.symbol[3])
			} else {
				steps += numeral.symbol
				result += numeral.symbol
			}
			n -= numeral.value
		}
	}
	return steps, result
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	input := os.Args[1]
	number := utils.BasicAtoi(input)
	if number <= 0 || number >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	steps, roman := intToRoman(number)
	fmt.Println(steps)
	fmt.Println(roman)
}
