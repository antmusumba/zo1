## HERE IS A SIMPLE PROGRAM TO TEST

package main

func main() {
mul := func(acc int, cur int) int {
return acc \* cur
}
sum := func(acc int, cur int) int {
return acc + cur
}
div := func(acc int, cur int) int {
return acc / cur
}
as := []int{500, 2}
ReduceInt(as, mul)
ReduceInt(as, sum)
ReduceInt(as, div)
}

# And its output :

$ go run .
1000
502
250
$
