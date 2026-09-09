package main

import "fmt"

func main() {
	// pascal1(5)
	m := pascal2(5)
	for _, v := range m {
		fmt.Println(v)
	}
}

func pascal1(n int) [][]int {
	m := [][]int{}

	if n == 1 {
		return append(m, []int{1})
	}
	if n == 2 {
		return append(m, []int{1}, []int{1, 1})
	}
	m = append(m, []int{1}, []int{1, 1})
	for i := 2; i < n; i++ {
		A := []int{1}
		for j := 0; j < len(m[i-1])-1; j++ {
			A = append(A, m[i-1][j]+m[i-1][j+1])
		}
		A = append(A, 1)

		m = append(m, A)
	}
	return m
}

func pascal2(n int) [][]int {
	m := [][]int{}
	if n == 1 {
		return append(m, []int{1})
	}
	if n == 2 {
		return append(m, []int{1}, []int{1, 1})
	}
	m = append(m, []int{1}, []int{1, 1})

	for i := 2; i < n; i++ {
		A := []int{1}
		for j := 1; j < i; j++ {
			A = append(A, m[i-1][j]+m[i-1][j-1])
		}
		A = append(A, 1)
		m = append(m, A)
	}
	return m
}

func ncr(n, r int) int64 {
	num, de := 1, 1
	for i := 0; i < r; i++ {
		num *= num - i
		de *= de - i
	}
	return int64(num / de)
}
