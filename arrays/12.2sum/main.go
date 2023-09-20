package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoSum_bruteForce([]int{2, 6, 5, 8, 11}, 14))
}

func twoSum_bruteForce(a []int, target int) (string, []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if target-a[i] == a[j] {
				return "yes", []int{i, j}
			}
		}
	}
	return "no", []int{}
}

func twoSum_better(a []int, target int) (string, []int) {
	m := map[int]int{}
	for i, v := range a {
		var j int
		var ok bool
		if j, ok = m[target-v]; ok {
			return "yes", []int{i, j}
		}
		m[v] = i
	}
	return "no", []int{}
}

func twoSum_optimal(a []int, target int) string {

	sort.Ints(a)
	l, h := 0, len(a)-1
	for l < h {
		sum := a[l] + a[h]
		if sum == target {
			return "yes"
		} else if sum > target {
			h--
		} else {
			l++
		}
	}
	return "no"
}
