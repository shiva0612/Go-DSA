package main

func main() {
	overlap([][]int{{1, 3}, {2, 6}, {8, 1}, {15, 8}})
}
func overlap(arr [][]int) [][]int {
	ans := [][]int{arr[0]}
	ci := 0

	for _, a := range arr {
		if ans[ci][1] >= a[0] {
			ans[ci][1] = a[1]
		} else {
			ci++
			ans = append(ans, a)
		}
	}
	return ans
}
