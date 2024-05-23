package utils

import (
	
	
	
)
func AlphaMirror(s string) string {
	mirror := ""
	for _, char := range s {
		mirror += string(char + 26)
	}
	return mirror

}