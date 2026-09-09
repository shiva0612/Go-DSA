package main

import (
	"math"
)

func main() {
	buysell([]int{7, 1, 5, 3, 6, 4})
}

func buysell_brute_force(d []int) int {
	maxProfit := 0
	for i := 0; i < len(d)-1; i++ {
		for j := i + 1; j < len(d); j++ {
			if d[j]-d[i] > maxProfit {
				maxProfit = d[j] - d[i]
			}
		}
	}
	return maxProfit
}

func buysell(d []int) int {
	min := math.MaxInt
	curP, maxp := 0, 0

	for i := 0; i < len(d); i++ {
		if d[i] < min {
			min = d[i]
		}
		curP = d[i] - min
		if curP > maxp {
			maxp = curP
		}
	}
	return maxp
}
