package main

func main() {
	search_in_2d([][]int{}, 2)
}

func search_in_2d(a [][]int, x int) bool {

	num := len(a) * len(a[0])
	l, h := 0, num-1
	for l <= h {
		mid := l + (h-l)/2
		r := mid / len(a)
		c := mid % len(a[0])
		if a[r][c] == x {
			return true
		}
		if a[r][c] > x {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false

}
