package utils

// import "github.com/01-edu/z01"
func Itoa(num int) string {
	result := ""
	neg := false

	if num < 0 {
		neg = true
		num = -num
	}
	if num == 0 {
		return "0"
	}

	for num > 0 {

		result = string(rune(num%10+'0')) + result
		num = num / 10
	}

	if neg {
		result = "-" + result
	}
	return result
}
