package main

import "fmt"

func main() {
	word := "   hello     world  hey    hhhh"
	str := plit(word)
	// fmt.Println(plit(str))
	res := ""
	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-1]{
			res += string(str[i]) + "   "
		} else {
			res += string(str[i])
		}
	}
	fmt.Println(res)

}

func plit(s string) []string {
	slc := []string{}
	result := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) != " " {
			result += string(s[i])
		} else if string(s[i]) == " " && result != "" {
			slc = append(slc, result)
			result = ""
		}
	}
	if result != "" {
		slc = append(slc, result)
	}

	return slc
}
