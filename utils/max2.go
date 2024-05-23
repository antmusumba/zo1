package utils

func Maxima(nb []int) int {
	for i := 0; i < len(nb); i++{
		if nb[i] > nb[0]{
			nb[0] = nb[i]
		}
	}
	return nb[0]
}