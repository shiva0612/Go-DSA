package main

import "fmt"

func main() {
	// a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	a := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	rotate_90(a)
	printArray(a)
}

func rotate_90(a [][]int) {

	for i := 0; i < len(a); i++ {
		for j := 0; j < i; j++ {
			a[i][j], a[j][i] = a[j][i], a[i][j]
		}
	}

	for i := 0; i < len(a); i++ {
		reverse(a[i])
	}
}

func reverse(v []int) {
	s, e := 0, len(v)-1
	for s < e {
		v[s], v[e] = v[e], v[s]
		s++
		e--
	}
}

func printArray(a [][]int) {
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
