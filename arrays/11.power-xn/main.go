package main

import "fmt"

func main() {
	ans := power(2, 10)
	fmt.Println(ans)
}
func power(x, n int) int {

	ans := 1
	for n > 0 {
		if n%2 == 1 {
			ans *= x
			n -= 1
		} else {
			x *= x
			n /= 2
		}
	}

	return ans
}
