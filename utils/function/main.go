package main

import "fmt"
func main(){
	//price:= mvuvi(2,30.8)
	fmt.Println(mvuvi(2,30.8))

}
func mvuvi(num int, price float64) float64 {
	totalPrice:= float64(num) * price 
	return totalPrice

}