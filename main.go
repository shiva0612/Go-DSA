package main

import (
	"fmt"
	"math"
)

func main() {
	ans := fruits([]int{1, 2, 1}, 2)
	fmt.Println(ans)
}
func fruits(a []int, baskets int) int {
	l := 0
	maxi := 0
	m := map[int]int{}
	for r := 0; r < len(a); r++ {
		m[a[r]] = r
		if len(m) > baskets {
			fruits := r - 1 - l + 1
			if fruits > maxi {
				maxi = fruits
			}

			//move L until map is again len=2
			index := math.MaxInt
			key := -1
			for k, v := range m {
				if v < index {
					index = v
					key = k
				}
			}
			l = m[key] + 1
			delete(m, key)

		}
	}
	return maxi
}
