package main

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
