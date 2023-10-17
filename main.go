package main

import (
	"fmt"
)

func main() {
	// a := []int{1}
	// ds := []int{1}
	// fmt.Println("initial : ", a, ds)
	// single4(&a, &ds)
	// fmt.Println("after func call: ", a, ds)
	// ds[0] = 2
	// fmt.Println("after changing :", a, ds)

	// a := []int{1}
	// ds := []int{1}
	// fmt.Println("initial : ", a, ds)
	// single5(&a, &ds)
	// fmt.Println("after func call: ", a, ds)
	// ds[0] = 2
	// fmt.Println("after changing :", a, ds)

	// a := []int{1}
	// ds := []int{1}
	// fmt.Println("initial : ", a, ds)
	// single6(&a, ds)
	// fmt.Println("after func call: ", a, ds)
	// ds[0] = 2
	// fmt.Println("after changing :", a, ds)

	a := []int{1}
	b := []int{0}
	copy(b, a)
	fmt.Println(a, b)

}

// this wont work
func single1(a []int) {
	a = append(a, 2)
}

// will work
func single2(a *[]int) {
	*a = append(*a, 2)
}

// wont work
func single3(a []int, ds []int) {
	a = append(a, ds...)
}

// will work - but need to check that ds change should not effect a
func single4(a *[]int, ds *[]int) {
	*a = append(*a, *ds...)
}

func single5(a *[]int, ds *[]int) {
	temp := []int{}
	copy(temp, *ds)
	*a = append(*a, temp...)
}
func single6(a *[]int, ds []int) {
	temp := []int{}
	copy(temp, ds)
	*a = append(*a, temp...)
}
