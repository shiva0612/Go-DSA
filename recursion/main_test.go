package main

import (
	"fmt"
	"testing"
)

func Test_sub_seq_sum(t *testing.T) {
	a := []int{1, 2, 3}
	sum := 3
	start := 0
	initial_sum := 0
	sub_seq_sum(start, len(a), initial_sum, sum, []int{}, a)
}
func Test_combination_sum(t *testing.T) {
	a := []int{1, 2, 3}
	sum := 3
	start := 0
	combination_sum(start, sum, []int{}, a)
}

func Test_combination_of_all_sums_in_subset(t *testing.T) {

	a := []int{3, 1, 2}
	ans := combination_of_all_sums_in_subset(0, 0, &[]int{}, a)
	fmt.Println(ans)
}
func Test_print_permunations(t *testing.T) {
	a := []int{1, 2, 3}
	ans := [][]int{}
	temp := []int{}
	m := map[int]bool{}
	ans = print_permunations(temp, a, m, ans)
	fmt.Println(ans)

}
