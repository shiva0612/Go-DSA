package main

import "fmt"

func main() {
	fmt.Println(brute_normal([]int{4, 12, 5, 3, 1, 2, 5, 3, 1, 2, 4, 6}))
	fmt.Println(brute_circular([]int{4, 12, 5, 3, 1, 2, 5, 3, 1, 2, 4, 6}))
	fmt.Println(brute_circular2([]int{4, 12, 5, 3, 1, 2, 5, 3, 1, 2, 4, 6}))

}
func brute_normal(a []int) []int {
	ng := []int{}

	for i := 0; i < len(a); i++ {
		found := false
		for j := i + 1; j < len(a); j++ {
			if a[j] > a[i] {
				ng = append(ng, a[j])
				found = true
				break
			}
		}
		if !found {
			ng = append(ng, -1)
		}
	}
	return ng
}
func brute_circular(a []int) []int {
	ng := []int{}

	for i := 0; i < len(a); i++ {
		found := false
		for j := i + 1; j < len(a); j++ {
			if a[j] > a[i] {
				ng = append(ng, a[j])
				found = true
				break
			}
		}
		if !found {
			for k := 0; k < i; k++ {
				if a[k] > a[i] {
					ng = append(ng, a[k])
					found = true
					break
				}
			}
		}
		if !found {
			ng = append(ng, -1)
		}
	}
	return ng
}
func brute_circular2(a []int) []int {
	ng := []int{}

	for i := 0; i < len(a); i++ {
		found := false
		for j := i + 1; j%len(a) != i; j++ {
			if j == len(a) {
				j = 0
			}
			if a[j] > a[i] {
				ng = append(ng, a[j])
				found = true
				break
			}
		}
		if !found {
			ng = append(ng, -1)
		}
	}
	return ng
}
