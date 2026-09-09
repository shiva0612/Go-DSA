package main

import "fmt"

func coins() {
	coins := []int{1, 2, 5, 10, 20, 50, 100}
	val := 87
	ans := []int{}
	count := 0

	for i := len(coins) - 1; i >= 0; i-- {

		if coins[i] <= val {
			val -= coins[i]
			ans = append(ans, coins[i])
			count += 1
			i += 1
		}
	}
	fmt.Println(ans)
}
