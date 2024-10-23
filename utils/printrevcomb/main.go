package main

import "github.com/01-edu/z01"

// 987, 986, 985, 984, 983, 982, 981, 980, 976, ..., 310, 210$
func main() {
	for i := '9'; i >= '2'; i-- {
		for j := i - 1; j >= '1'; j-- {
			for k := j - 1; k >= '0'; k-- {
				if i == '2' && j == '1' && k == '0' {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
				} else {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)

					z01.PrintRune(',')
					z01.PrintRune(' ')

				}
			}
		}
	}
	z01.PrintRune('\n')
}
