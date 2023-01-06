package main

import "fmt"

func main() {
	kadanes([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
}

func kadanes(a []int) int {
	msum, csum := 0, 0
	s := 0
	ans := []int{}

	for i := 0; i < len(a); i++ {
		csum += a[i]
		if csum > msum {
			msum = csum
			ans = []int{s, i}
		}
		if csum < 0 {
			s = i + 1
			csum = 0
		}
	}
	fmt.Println(a[ans[0] : ans[1]+1])
	sum := 0
	for _, v := range a[ans[0] : ans[1]+1] {
		sum += v
	}
	return sum
}
