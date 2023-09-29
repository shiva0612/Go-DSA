package main

import "math"

// upper bound
func sq_root(n int) int {
	l, h := 0, n
	ans := 1
	for l <= h {
		m := (l + h) / 2
		if m*m <= n {
			ans = m
			l = m + 1
		} else {
			h = m - 1
		}
	}
	return ans
}

func nthRoot(n, m int) int {
	nroot := func(mid, n, m int) int {
		ans := 1
		for i := 0; i < n; i++ {
			ans *= mid
			if ans > m {
				return 2
			}
		}
		if ans == m {
			return 1
		}
		return 0
	}

	l, h := 0, m
	for l <= h {
		mid := (l + h) / 2
		mroot := nroot(mid, n, m)
		if mroot == 1 {
			return mid
		}
		if mroot == 2 {
			//greater
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func minRateToEatBananas(a []int, hours int) int {
	hoursRequired := func(a []int, hourly int) int {
		total := 0
		for i := 0; i < len(a); i++ {
			q := a[i] / hourly
			if a[i]%hourly != 0 {
				q += 1
			}
			total += q
		}
		return total
	}

	l, h := 0, a[len(a)-1]
	ans := math.MaxInt
	for l <= h {
		m := (l + h) / 2
		computed := hoursRequired(a, m)
		if computed <= hours {
			h = m - 1
			if m < ans {
				ans = m
			}
		} else {
			l = m + 1
		}
	}
	return ans
}

func noOfBoq(a []int, boq, adj int) int {

	possible := func(a []int, day, boq, adj int) bool {

		possibleBoq := 0
		count := 0
		for _, v := range a {
			if day >= v {
				count += 1
			} else {
				count /= adj
				possibleBoq += count
				count = 0
			}
		}
		possibleBoq += count / adj
		return possibleBoq >= boq
	}

	if len(a) < boq*adj {
		return -1
	}

	mini := math.MaxInt
	maxi := math.MinInt
	for _, v := range a {
		if v <= mini {
			mini = v
		}
		if v >= maxi {
			maxi = v
		}
	}

	for i := mini; i < maxi+1; i++ {
		if possible(a, i, boq, adj) {
			return i
		}
	}
	return -1
}

func noOfBoq_bs(a []int, boq, adj int) int {

	possible := func(a []int, day, boq, adj int) bool {

		possibleBoq := 0
		count := 0
		for _, v := range a {
			if day >= v {
				count += 1
			} else {
				count /= adj
				possibleBoq += count
				count = 0
			}
		}
		possibleBoq += count / adj
		return possibleBoq >= boq
	}

	if len(a) < boq*adj {
		return -1
	}

	mini := math.MaxInt
	maxi := math.MinInt
	for _, v := range a {
		if v <= mini {
			mini = v
		}
		if v >= maxi {
			maxi = v
		}
	}

	l, h := mini, maxi
	ans := math.MaxInt
	for l <= h {
		mid := (l + h) / 2
		if possible(a, mid, boq, adj) {
			h = mid - 1
			if mid < ans {
				ans = mid
			}
		} else {
			l = mid + 1
		}
	}
	return ans
}

func min_ship_days(a []int, maxDays int) int {

	daysReq := func(a []int, cap int) int {
		days := 1
		load := 0

		for _, v := range a {
			if load+v > cap {
				days += 1
				load = v
			} else {
				load += v
			}
		}

		return days
	}

	maxi := math.MinInt
	sum := 0
	for _, v := range a {
		sum += v
		if v >= maxi {
			maxi = v
		}
	}

	l, h := maxi, sum
	ans := math.MaxInt
	for l <= h {
		mid := (l + h) / 2

		if daysReq(a, mid) <= maxDays {
			h = mid - 1
			if mid < ans {
				ans = mid
			}
		} else {
			l = mid + 1
		}
	}
	return ans
}




