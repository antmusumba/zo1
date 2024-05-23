package utils

//import "github.com/01-edu/z01"

func Itoa(num int)string{
	result := ""
	for num != 0{
		mod := num % 10
		startrune := '0'
		for i := 0; i < mod; i++{
			startrune++
		}
		result = result + string(startrune)
		num = num / 10
	}
	result2 :=""
	for j:= len(result)-1;j >= 0; j-- {
		result2 += string(result[j])
	}
	// for _, val := range result2 {
	// 	z01.PrintRune(val)
	// }
	// z01.PrintRune('\n')
	return result2
}