package main

import "fmt"

func main() {
	missing, repeating := missing_repeating([]int{3, 1, 2, 5, 3})
	fmt.Println("missing = ", missing)
	fmt.Println("repeating = ", repeating)
}

func missing_repeating(a []int) (missing, repeating int) {

	xr := 0
	for i := 0; i < len(a); i++ {
		xr ^= a[i]
		xr ^= i + 1
	}

	bitNumber := 0
	for {
		if (xr & (1 << bitNumber)) != 0 {
			break
		}
		bitNumber++
	}

	zeroClub := 0
	oneClub := 0
	for i := 0; i < len(a); i++ {
		if (a[i] & (1 << bitNumber)) != 0 {
			oneClub ^= a[i]
		} else {
			zeroClub ^= a[i]
		}
	}
	for i := 1; i < len(a)+1; i++ {
		if (i & (1 << bitNumber)) != 0 {
			oneClub ^= i
		} else {
			zeroClub ^= i
		}
	}

	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] == zeroClub {
			count++
		}
	}

	if count == 2 {
		return oneClub, zeroClub
	}
	return zeroClub, oneClub
}
