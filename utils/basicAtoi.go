package utils

//"fmt"

//"github.com/01-edu/z01"

func BasicAtoi(s string) int {
	if s == "0" {
		return 0
	}
	neg := false
	if s[0] == '-' {
		neg = true
		s = s[1:]
	}

	y := 0
	for _, val := range s {
		y = y*10 + int(val-'0')
	}
	if neg {
		y = -y
	}
	return y
}

