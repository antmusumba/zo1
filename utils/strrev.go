package utils


func StrRev(s string) string{
	var y string
	for _, i := range s {
		y = string(i) + y
	}
	return y
}