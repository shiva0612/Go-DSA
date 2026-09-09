package main

import "fmt"

func main() {
	a := overlap([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 8}})
	fmt.Println(a)
}
func overlap(arr [][]int) [][]int {
	ans := [][]int{arr[0]}

	for _, a := range arr[1:] {
		ci := len(ans) - 1
		if a[0] < ans[ci][1] {
			ans[ci][1] = a[1]
		} else {
			ans = append(ans, a)
		}
	}
	return ans
}
