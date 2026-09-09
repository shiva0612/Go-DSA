package main

import (
	"fmt"
	"math"
)

func main() {

}

func find_peak_bs(a []int) int {
	l, h := 0, len(a)-1
	if l == h {
		return a[l]
	}
	if a[l] > a[l+1] {
		return a[l]
	}
	if a[h] > a[h-1] {
		return a[h]
	}
	l += 1
	h -= 1
	for l <= h {
		m := l + h/2
		if a[m] > a[m-1] && a[m] > a[m+1] {
			return a[m]
		}
		//elimination
		if a[m] > a[m-1] {
			//increasing curve => right side peak will be there
			l = m + 1
		} else {
			//decreasing curve => left side peak will be there
			h = m - 1
		}
	}
	return -1

}

// considering -inf [arr...] -inf
func find_peak_linear(a []int) int {
	for i := 0; i < len(a); i++ {
		if i == 0 && a[i] > a[i+1] {
			return i
		}
		if i == len(a)-1 && a[i] > a[i-1] {
			return i
		}
		if a[i] > a[i-1] && a[i] > a[i+1] {
			return a[i]
		}
	}
	return -1
}

func peak(a []int) {
	for i := 0; i < len(a); i++ {
		if i == 0 && a[i] > a[i+1] {
			fmt.Println(a[i])
		}
		if i == len(a)-1 && a[i] > a[i-1] {
			fmt.Println(a[i])
		}
		if a[i] > a[i-1] && a[i] > a[i+1] {
			fmt.Println(a[i])
		}
		if a[i] < a[i-1] && a[i] < a[i+1] {
			fmt.Println(a[i])
		}
	}
}

func bs_iterative(a []int, target int) int {
	l, h := 0, len(a)
	for l <= h {
		m := l + h/2
		if a[m] == target {
			return m
		}
		if a[m] > target {
			h = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

func bs_recursive(a []int, l, h, target int) int {
	if l > h {
		return -1
	}
	m := l + h/2
	if a[m] == target {
		return m
	} else if a[m] > target {
		return bs_recursive(a, 0, m-1, target)
	}
	return bs_recursive(a, m+1, h, target)
}

func find_single_element(a []int) int {
	l, h := 0, len(a)-1
	if l == h {
		return a[l]
	}
	if a[0] != a[1] {
		return a[0]
	}
	if a[h] != a[h-1] {
		return a[h]
	}
	l += 1
	h -= 1
	for l <= h {
		m := (l + h) / 2
		if a[m] != a[m+1] && a[m] != a[m-1] {
			return a[m]
		}
		//elemination
		if (m%2 == 1 && a[m] == a[m-1]) ||
			(m%2 == 0 && a[m] == a[m+1]) {
			//we are on left of ans - move to right
			l = m + 1
		} else {
			//we are on right of ans - move to left
			h = m - 1
		}
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func point_of_rotation(a []int) {
	l, h := 0, len(a)-1
	ans := math.MaxInt

	for l <= h {
		m := (l + h) / 2
		if a[l] < a[m] {
			//left is sorted
			//update the min in ans by taking the first element which is a[l]
			ans = min(ans, a[l])

			l = m + 1
		} else {
			//right half is sorted
			//update the min in ans by taking the first element which is a[m]
			ans = min(ans, a[m])

			h = m - 1
		}
	}
}

func times_rotated(a []int) {
	l, h := 0, len(a)-1
	ans := math.MaxInt
	index := -1
	for l <= h {
		m := (l + h) / 2
		if a[l] < a[m] {
			//left is sorted
			//update the min in ans by taking the first element which is a[l]
			if ans < a[l] {
				ans = a[l]
				index = l
			}

			l = m + 1
		} else {
			//right half is sorted
			//update the min in ans by taking the first element which is a[m]
			if ans < a[m] {
				ans = a[m]
				index = m
			}

			h = m - 1
		}
	}
	fmt.Println(index)
}

func search_in_rotated_with_duplicate(a []int, k int) int {
	l, h := 0, len(a)-1
	for l <= h {
		m := (l + h) / 2
		if a[m] == k {
			return k
		}

		if a[l] == a[m] && a[m] == a[h] {
			l += 1
			h -= 1
			continue
		}
		if a[l] <= a[m] {
			//left is sorted
			if a[l] <= k && k <= a[m] {
				//l k m
				h = m - 1
			} else {
				l = m + 1
			}
		} else {
			//right is sorted
			if a[m] <= k && k <= a[h] {
				//m k h
				l = m + 1
			} else {
				h = m - 1
			}
		}
	}
	return -1
}

func search_in_rotated(a []int, k int) int {
	l, h := 0, len(a)-1
	for l <= h {
		m := (l + h) / 2
		if a[m] == k {
			return k
		}
		if a[l] <= a[m] {
			//left is sorted
			if a[l] <= k && k <= a[m] {
				//l k m
				h = m - 1
			} else {
				l = m + 1
			}
		} else {
			//right is sorted
			if a[m] <= k && k <= a[h] {
				//m k h
				l = m + 1
			} else {
				h = m - 1
			}
		}
	}
	return -1
}

func first_last_occurance(a []int, k int) []int {
	return []int{lowerBound(a, k), upperBound(a, k)}
}

func count_occurances(a []int, k int) int {
	return upperBound(a, k) - lowerBound(a, k) + 1
}

// ciel
func lowerBound(a []int, x int) int {
	l, h := 0, len(a)-1
	ans := len(a)
	for l <= h {
		m := (l + h) / 2
		if a[m] >= x {
			ans = m
			h = m - 1
		} else {
			l = m + 1
		}
	}
	return ans
}

// floor
func upperBound(a []int, x int) int {
	l, h := 0, len(a)-1
	ans := len(a)
	for l <= h {
		m := (l + h) / 2
		if a[m] > x {
			ans = m
			l = m + 1
		} else {
			h = m - 1
		}
	}
	return ans
}
