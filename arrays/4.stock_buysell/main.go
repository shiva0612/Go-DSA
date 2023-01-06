package main

import (
	"fmt"
	"math"
)

func main() {
	buysell([]int{7, 1, 5, 3, 6, 4})
}

func buysell(d []int) int {
	mini, maxi := 0, 0
	min, max := math.MaxInt, 0
	curP, maxp := 0, 0

	for i := 0; i < len(d); i++ {
		if d[i] < min {
			mini, min = i, d[i]
		}
		curP = d[i] - min
		if curP > maxp {
			maxi, max = i, d[i]
			maxp = curP
		}
	}
	fmt.Printf("buy at d[%d]=%d & sell at d[%d]=%d = profit[%d]", mini, min, maxi, max, maxp)
	return maxp
}
