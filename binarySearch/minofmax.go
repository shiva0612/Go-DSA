package main

import "fmt"

func main() {
	// fmt.Println(aggCow_bs([]int{0, 3, 4, 7, 9, 10}, 4))
	// fmt.Println(aggCow([]int{0, 3, 4, 7, 9, 10}, 4))
}

func possible(a []int, dist, n int) bool {
	last := a[0]
	count := 1
	for i := 1; i < len(a); i++ {
		if a[i]-last >= dist {
			count++
			last = a[i]
		}
		if count >= n {
			return true
		}
	}
	return false
}

func aggCow(a []int, n int) int {

	l := len(a)
	min, max := a[0], a[l-1]
	ans := -1
	for i := 1; i < max-min+1; i++ {
		if !possible(a, i, n) {
			break
		}
		ans = i
	}
	return ans
}

func aggCow_bs(a []int, n int) int {

	l, h := 1, len(a)-1
	ans := -1
	for l <= h {
		mid := (l + h) / 2
		if possible(a, mid, n) {
			//since we want max, we move right
			l = mid + 1
			ans = mid
		} else {
			h = mid - 1
		}
	}
	return ans
}
