package main

import "fmt"

func ReverseBits(oct byte) byte {
	var result byte = 0 

	for i := 0 ; i < 8 ; i++ {
		result = (result << 1) | (oct & 1) // Shift result to the left and add the least significant bit of oct
		oct >>= 1
	}
	return result
}
func main(){
	oct := byte(0b00100110)
	reversed := ReverseBits(oct)
	fmt.Printf("Original: %08b, Reversed: %08b\n", oct, reversed) 
}