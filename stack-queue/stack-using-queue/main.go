package main

import "fmt"

type Q []int

var q1, q2 Q

func (q *Q) push(n int) {
	*q = append([]int{n}, *q...)
}
func (q *Q) pop() int {
	n := (*q)[0]
	if len(*q) > 1 {
		*q = (*q)[1:]
		return n
	}
	*q = []int{}
	return n
}
func (q Q) get() []int {
	return q
}

func main() {
	push(1)
	push(2)
	push(3)
	push(4)
	push(5)
	pop()
	printStack()
}
func push(n int) {
	q2.push(n)
	for len(q1) > 0 {
		q2.push(q1.pop())
	}
	q1, q2 = q2, q1
}
func pop() {
	q1.pop()
}
func printStack() {
	fmt.Println(q1.get())
}
