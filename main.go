package main

import "zo1/utils"

//"os"
// "zo1/utils"

// "fmt"

//"github.com/01-edu/z01"
// "golang.org/x/tools/go/analysis/passes/stringintconv"

//"os"

//"github.com/01-edu/z01"
//"zo1/utils"
//"github.com/01-edu/z01"
// "fmt"
// "zo1/utils"
// "github.com/01-edu/z01"
// "fmt"
// "os"
// "fmt"
// "fmt"
// "zo1/utils"
// "zo1/utils"
// "github.com/01-edu/z01"

// func main() {
// utils.AbCd()

// a := []int{23, 123, 1, 11, 55, 93}
// max := utils.Max(a)

// fmt.Println(max)
// // utils.DescendComb2()
// args := os.Args[1]
// // utils.StrLen(args)
// utils.PrintNbr(-123)
// utils.PrintNbr(0)
// utils.PrintNbr(123)
// z01.PrintRune('\n')
// myArgs := os.Args[1:]
// count := 0
// for i := 0 ; i < len(myArgs) ; i++ {
// 	count++
// }
// fmt.Println(count)
// utils.PrintHello()
// fmt.Println(utils.BasicAtoi("34"))

// fmt.Println('1' - 48)
// dat := utils.AlphaMirror("HELLO")
// fmt.Println(dat)
// utils.PrintAlphabets()
// dat := utils.BasicAtoi("56")
// var r []rune
// for dat > 0 {
// 	num := dat % 10
// 	dat /= 10
// 	r = append(r, rune(num+'0'))
// }
// for i := len(r) - 1; i >= 0; i-- {
// 	z01.PrintRune(r[i])
// }
// z01.PrintRune('\n')

// utils.PrintNbr(998765439)
// dat := utils.StrRev("hello")
// for _, i := range dat {
// 	z01.PrintRune(i)
// }
// //z01.PrintRune('\n')
// dat := utils.Itoa(78)
// fmt.Println(dat)
// 	count := 0
// 	for i := 0; i < len(os.Args[1:]); i++ {
// 		count++
// 	}
// 	z01.PrintRune(rune(count) - '0')
// 	z01.PrintRune('\n')
// }
/*func main(){
	args := os.Args[1:]

	count := 0
	for range args{
		count++
	}

	result := ""
	for count != 0{
		mod := count % 10
		startrune := '0'

		for  i := 0 ; i < mod ; i++{
			startrune++
		}
		result = string(startrune) + result
		count = count / 10
	}

	for _, chr := range result{
		z01.PrintRune(chr)
	}
	z01.PrintRune('\n')
}
*/
/*
	func main(){
		args := os.Args[1:]
		count := 0
		for range args {
			count++

		}
		//fmt.Println(count-'0')
		z01.PrintRune(rune(count-'0'))



		res := ""

		for count != 0 {
			mod := count % 10
			startrune := '0'
			for i := 0 ; i < mod ; i++ {
				startrune++
			}
			res =string(startrune) + res
			count= count/10

		}
		for _, ch := range res {
			z01.PrintRune(ch)
		}

		z01.PrintRune('\n')
		//utils.Itoa(2345)




	}
*/
/*func main(){
if len(os.Args) != 3 {
	return
}
str1:=os.Args[1]
str2:=os.Args[2]

result := ""
for _, ch := range str1{
	outcome := false
	for i, chr := range str2 {
		if ch == chr {
			result += string(chr)
			outcome = true

			str2 = str2[i:]
			break
		}
	}
	if !outcome {
		return
	}
}
for _, ch := range result {
	z01.PrintRune(ch)
}
z01.PrintRune('\n')
}
*/
		func main(){
			//utils.Itoa(56)
			utils.PrintNbr(-467)
		}


	
 
