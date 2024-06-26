package main

import "fmt"

// SwapBits swaps the halves of a byte.
func SwapBits(octet byte) byte {
	// Right shift the first 4 bits to the last 4 positions
	leftHalf := octet >> 4
	// Left shift the last 4 bits to the first 4 positions
	rightHalf := octet << 4
	// Combine both halves
	return leftHalf | rightHalf
}

func main() {
	// Example usage
	octet := byte(0x41) // 0100 0001 in binary
	swapped := SwapBits(octet)
	fmt.Printf("Original byte: %08b\n", octet)
	fmt.Printf("Swapped byte:  %08b\n", swapped)
}
