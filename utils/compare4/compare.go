package main

import(
	"fmt"
)
func main () {

	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "slut!"))
	fmt.Println(Compare("Ola!", "Oladfdgfg"))






}
func Compare(a, b string) int {

if a== b {
	return 0
} else if a >b {
	return 1
} else {
	return -1
}

}

