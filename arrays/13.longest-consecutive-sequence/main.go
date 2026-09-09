package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	fmt.Println(brute([]int{1, 1, 1, 2, 2, 100, 101, 102, 7, 8}))
	fmt.Println(better([]int{1, 1, 1, 2, 2, 100, 101, 102, 7, 8}))
	fmt.Println(optimal([]int{1, 1, 1, 2, 2, 100, 101, 102, 7, 8}))

}

func brute(a []int) int {

	nxtNumberPresent := func(a []int, target int) bool {
		for _, v := range a {
			if v == target {
				return true
			}
		}
		return false
	}

	longest := 0
	for i := 0; i < len(a); i++ {
		x := a[i]
		count := 1
		for nxtNumberPresent(a, x+1) {
			count++
			x++
		}
		if count > longest {
			longest = count
		}
	}
	return longest
}

func better(a []int) int {

	sort.Ints(a)
	lastSmallestFound := math.MinInt32
	var count, longest int
	for _, v := range a {
		if v-1 == lastSmallestFound {
			count++
			lastSmallestFound = v
		} else if v != lastSmallestFound {
			lastSmallestFound = v
			count = 1
		}
		if count > longest {
			longest = count
		}
	}
	return longest
}

func optimal(a []int) int {

	m := map[int]bool{}
	for _, v := range a {
		m[v] = true
	}

	longest := 1

	for k, _ := range m {
		if !m[k-1] {
			count := 1
			x := k
			for m[x+1] {
				x++
				count++
			}
			if count > longest {
				longest = count
			}
		}
	}

	return longest

}
