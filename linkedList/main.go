package main

import (
	"fmt"
)

type List struct {
	Val int
	Nxt *List
}

func NewList(v int) *List {
	return &List{Val: v}
}

func NewListFromArray(a []int) *List {

	if len(a) <= 0 {
		return nil
	}
	head := NewList(a[0])
	temp := head
	for _, v := range a[1:] {
		temp.Nxt = NewList(v)
		temp = temp.Nxt
	}

	return head
}
func (l *List) print() {
	temp := l
	fmt.Println("")
	for temp != nil {
		fmt.Print(temp.Val, " ")
		temp = temp.Nxt
	}
	fmt.Println("")
}

func (l *List) arr() []int {
	t := l
	a := []int{}
	for t != nil {
		a = append(a, t.Val)
		t = t.Nxt
	}
	return a
}

func (l *List) find(n int) *List {
	t := l
	for t != nil {
		if t.Val == n {
			return t
		}
		t = t.Nxt
	}
	return nil
}
func main() {

	p := NewListFromArray([]int{1, 3, 1, 2, 4})
	q := NewListFromArray([]int{3})
	q.Nxt = p.find(2)
	fmt.Println(findIntersection(p, q).Val)
}

func findIntersection(p, q *List) *List {

	//assuming that intersection is there for sure
	x, y := p, q
	for x != y {
		x = x.Nxt
		y = y.Nxt
		if y == nil {
			y = p
		}
		if x == nil {
			x = q
		}
	}
	return x
}

func isLoop(head *List) bool {

	s, f := head, head
	for f.Nxt != nil {

		s = s.Nxt
		f = f.Nxt.Nxt
	}
	return s == f
}

func loopNode(head *List) *List {

	s, f, e := head, head, head
	for f.Nxt != nil {

		s = s.Nxt
		f = f.Nxt.Nxt
	}
	if s != f {
		return nil
	}
	for e != s {
		e = e.Nxt
		s = s.Nxt
	}
	return s

}
