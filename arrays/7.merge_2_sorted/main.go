package main

import (
	"fmt"
	"sort"
)

func main() {
	merge2sorted([]int{1, 4, 7, 8, 10}, []int{2, 3, 9})
}

func merge2sorted(a []int, b []int) {
	for i := 0; i < len(a); i++ {
		if a[i] > b[0] {
			a[i], b[0] = b[0], a[i]
			sort.Ints(b)
		}
	}
	fmt.Println(a, b)
}
