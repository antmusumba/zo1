package main

import "fmt"

func Max(a []int) int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a[len(a)-1]
}
func main(){
	n := []int{1,56,3,4,5,6,7,}
	max := Max(n)
	fmt.Println(max)

}
