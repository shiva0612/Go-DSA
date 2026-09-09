package main

import "fmt"

func nextGreater(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	// Stack to store elements
	stack := []int{}

	// Traverse from right to left
	for i := n - 1; i >= 0; i-- {

		// Remove all smaller or equal elements
		for len(stack) > 0 && nums[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		// If stack is empty, no greater element
		if len(stack) == 0 {
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-1]
		}

		// Push current element
		stack = append(stack, nums[i])
	}

	return res
}

func main() {
	nums := []int{4, 5, 2, 10}

	ans := nextGreater(nums)

	fmt.Println(ans)
}
