package main

import "fmt"

func main() {
	rotate_90([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}

func rotate_90(a [][]int) {
	for i := 1; i < len(a); i++ {
		a[i][i-1], a[i-1][i] = a[i-1][i], a[i][i-1]
	}
	a[0][len(a)-1], a[len(a)-1][0] = a[len(a)-1][0], a[0][len(a)-1]
	for _, v := range a {
		reverse(v)
	}
	fmt.Println(a)
}
func reverse(v []int) {
	s, e := 0, len(v)-1
	for s < e {
		v[s], v[e] = v[e], v[s]
		s++
		e--
	}
}
