package main

import "math"

func bs_2d(a [][]int, target int) bool {

	n := len(a)
	m := len(a[0])
	l, h := 0, n*m-1
	for l <= h {
		mid := (l + h) / 2
		row := mid / m
		col := mid % m
		v := a[row][col]
		if v == target {
			return true
		} else if v < target {
			l = mid + 1
		} else {
			h = mid - 1
		}
	}
	return false
}

func median_of_sorted_bs(a, b []int) float32 {

	maxi := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	mini := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	n1, n2 := len(a), len(b)
	if n1 > n2 {
		return median_of_sorted_bs(b, a)
	}

	n := n1 + n2
	left := (n + 1) / 2
	l, h := 0, n1
	for l <= h {
		m1 := (l + h) / 2
		m2 := left - m1

		l1 := math.MinInt
		l2 := l1
		r1 := math.MaxInt
		r2 := r1

		if m1-1 >= 0 {
			l1 = a[m1-1]
		}
		if m2-1 >= 0 {
			l2 = b[m2-1]
		}
		if m1 < n1 {
			r1 = a[m1]
		}
		if m2 < n2 {
			r2 = a[m2]
		}

		if l1 <= r2 && l2 <= r1 {
			if n%2 == 1 {
				return float32(maxi(l1, l2))
			} else {
				return float32(maxi(l1, l2)+mini(r1, r2)) / 2.0
			}
		}
		if l1 > r2 {
			h = m1 - 1
		} else {
			l = m1 + 1
		}
	}
	return 0
}

func median_of_sorted(a, b []int) float32 {

	i, j := 0, 0
	n1, n2 := len(a), len(b)
	n := n1 + n2
	id2 := n / 2
	id1 := id2 - 1
	idx := 0
	e1, e2 := -1, -1
	for i < n1 && j < n2 {
		if a[i] < b[i] {
			if i == id1 {
				e1 = a[i]
			}
			if i == id2 {
				e2 = a[i]
			}
			i++
		} else {
			if i == id1 {
				e1 = b[i]
			}
			if i == id2 {
				e2 = b[i]
			}
			j++
		}
		idx++
	}
	if i < n1 {
		if i == id1 {
			e1 = a[i]
		}
		if i == id2 {
			e2 = a[i]
		}
		i++
		idx++
	}
	if j < n2 {
		if i == id1 {
			e1 = b[i]
		}
		if i == id2 {
			e2 = b[i]
		}
		j++
		idx++
	}
	if n%2 == 1 {
		return float32(e2)
	}
	return float32((e1 + e2)) / 2.0
}

func possible(a []int, dist, n int) bool {
	last := a[0]
	count := 1
	for i := 1; i < len(a); i++ {
		if a[i]-last >= dist {
			count++
			last = a[i]
		}
		if count >= n {
			return true
		}
	}
	return false
}

func aggCow(a []int, n int) int {

	l := len(a)
	min, max := a[0], a[l-1]
	ans := -1
	for i := 1; i < max-min+1; i++ {
		if !possible(a, i, n) {
			break
		}
		ans = i
	}
	return ans
}

func aggCow_bs(a []int, n int) int {

	l, h := 1, len(a)-1
	ans := -1
	for l <= h {
		mid := (l + h) / 2
		if possible(a, mid, n) {
			//since we want max, we move right
			l = mid + 1
			ans = mid
		} else {
			h = mid - 1
		}
	}
	return ans
}
