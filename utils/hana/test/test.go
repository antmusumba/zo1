package main

import "fmt"

func main() {
	word := " helo world of code    "
	str := splitThat(word, " ")
	res := ""
	s := str[0]
	fmt.Println(">>", string(s[len(s)-1]-32))

	b := []byte(word)

	for i := 0; i < len(b); i++ {
		if i+1 < len(b) {
			if b[i+1] == ' ' {
				b[i] = b[i] - 32
			}
		}
	}
	fmt.Println(">>", string(b))

	fmt.Println(res)
}

func splitThat(s string, sep string) []string {
	slc := []string{}
	result := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) != sep {
			result += string(s[i])
		} else if string(s[i]) == sep && result != "" {
			slc = append(slc, result)
			result = ""
		}
	}
	if result != "" {
		slc = append(slc, result)
	}
	return slc
}
