package main

import (
	"fmt"
)

// ------------------------------------------------------------------------------------
func print_permunations(temp, a []int, m map[int]bool, ans [][]int) [][]int {

	if len(temp) == len(a) {
		if len(temp) == len(a) {
			ans = append(ans, temp)
		}
		return ans
	}

	for i := 0; i < len(a); i++ {
		picked, _ := m[i]
		if picked {
			continue
		}
		temp = append(temp, a[i])
		m[i] = true
		ans = print_permunations(temp, a, m, ans)

		temp = temp[:len(temp)-1]
		m[i] = false
	}
	return ans
}

// ------------------------------------------------------------------------------------
// all sums possible with in an array
// IMP : append creates new array hence
func combination_of_all_sums_in_subset(i, csum int, ans *[]int, a []int) []int {

	if i >= len(a) {
		*ans = append(*ans, csum)
		fmt.Println(ans)
		return *ans
	}
	combination_of_all_sums_in_subset(i+1, csum+a[i], ans, a)
	combination_of_all_sums_in_subset(i+1, csum, ans, a)
	return *ans
}
func combination_of_all_sums_in_subset2(i, csum int, ans []int, a []int) []int {

	if i >= len(a) {
		ans = append(ans, csum)
		fmt.Println(ans)
		return ans
	}
	ans = combination_of_all_sums_in_subset2(i+1, csum+a[i], ans, a)
	ans = combination_of_all_sums_in_subset2(i+1, csum, ans, a)
	return ans
}

func combination_of_all_sums_in_subset_wont_work(i, csum int, ans []int, a []int) []int {

	if i >= len(a) {
		ans = append(ans, csum)
		fmt.Println(ans)
		return ans
	}
	combination_of_all_sums_in_subset2(i+1, csum+a[i], ans, a)
	combination_of_all_sums_in_subset2(i+1, csum, ans, a)
	return ans
}

// ------------------------------------------------------------------------------------
// same element can be picked multiple times
func combination_sum(i, sum int, ans, a []int) {
	if i >= len(a) {
		if sum == 0 {
			fmt.Println(ans)
		}
		return
	}
	if sum >= a[i] {
		ans = append(ans, a[i])
		combination_sum(i, sum-a[i], ans, a) //pick
		ans = ans[:len(ans)-1]
	}
	combination_sum(i+1, sum, ans, a) //not-pick
}

// ------------------------------------------------------------------------------------

func sub_seq_sum(i, n, csum, sum int, ans, a []int) {
	if i == n {
		if csum == sum {
			fmt.Println(ans)
		}
		return
	}
	ans = append(ans, a[i])
	sub_seq_sum(i+1, n, sum_array(ans), sum, ans, a)
	ans = ans[:len(ans)-1]
	sub_seq_sum(i+1, n, sum_array(ans), sum, ans, a)
}

func sum_array(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

// ------------------------------------------------------------------------------------
/*
a := []int{5, 2, 8, 1, 3, 7, 6, 4}
mergeSort(a, 0, len(a)-1)
fmt.Println(a)
*/
func mergeSort(a []int, l, h int) {
	if l >= h {
		return
	}
	mid := l + (h-l)/2
	mergeSort(a, l, mid)
	mergeSort(a, mid+1, h)
	merge(a, l, mid, h)
}

func merge(a []int, l, mid, h int) {
	temp := make([]int, 0, h-l+1)
	i, j := l, mid+1
	for i <= mid && j <= h {
		if a[i] <= a[j] {
			temp = append(temp, a[i])
			i++
		} else {
			temp = append(temp, a[j])
			j++
		}
	}
	for i <= mid {
		temp = append(temp, a[i])
		i++
	}
	for j <= h {
		temp = append(temp, a[j])
		j++
	}
	copy(a[l:h+1], temp)
}

// ------------------------------------------------------------------------------------
/*
take or not_take
a := []int{1, 2, 3}
generate_all_sub_seq(0, len(a), []int{}, a)
*/
func generate_all_sub_seq(i, n int, ans, a []int) {
	if i == n {
		fmt.Println(ans)
		return
	}
	ans = append(ans, a[i])
	generate_all_sub_seq(i+1, n, ans, a)
	ans = ans[:len(ans)-1]
	generate_all_sub_seq(i+1, n, ans, a)
}

// ------------------------------------------------------------------------------------
/*
ans := fibo(7)
fmt.Println(ans)
*/
func fibo(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}

// ------------------------------------------------------------------------------------
/*
a := []int{1, 2, 2, 1}
ans := check_palindrome(0, len(a)-1, a)
fmt.Println(ans)
*/
func check_palindrome(l, h int, a []int) bool {
	if l > h {
		return true
	}
	if a[l] == a[h] {
		return check_palindrome(l+1, h-1, a)
	}
	return false
}

// ------------------------------------------------------------------------------------
// reverse_array(0, 4, []int{1, 2, 3, 4, 5})
func reverse_array(l, h int, a []int) []int {
	if l > h {
		fmt.Println(a)
		return a
	}
	a[l], a[h] = a[h], a[l]
	return reverse_array(l+1, h-1, a)
}

// ------------------------------------------------------------------------------------
// n_sum(3, 0)
func n_sum(n, sum int) int {
	if n <= 0 {
		fmt.Println(sum)
		return sum
	}
	return n_sum(n-1, sum+n)
}

// ------------------------------------------------------------------------------------
func print_names_n_times(n int) {
	if n <= 0 {
		return
	}
	print_names_n_times(n - 1)
	fmt.Println("shiva: ", n)
}

func print_names_n_times_reverse_1(n int) {
	if n <= 0 {
		return
	}
	fmt.Println("shiva: ", n)
	print_names_n_times(n - 1)
}

func print_names_n_times_reverse_2(i, n int) {
	if i >= n {
		return
	}
	print_names_n_times_reverse_2(i+1, n)
	fmt.Println("shiva: ", i)
}
