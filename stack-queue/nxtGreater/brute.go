package main

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
