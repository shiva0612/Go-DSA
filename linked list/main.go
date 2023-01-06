package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	val  int
	next *Node
}

type List struct {
	head *Node
}

func mergeList(a, b *List) *List {
	l := new(List)

	i, j, k := a.head, b.head, l.head

	for i != nil && j != nil {
		node := new(Node)

		if i.val <= j.val {
			node.val = i.val
			i = i.next
		} else {
			node.val = j.val
			j = j.next
		}
		if l.head == nil {
			l.head = node
			k = l.head
		} else {
			k.next = node
			k = k.next
		}
	}
	if i == nil {
		for j != nil {
			k.next = j
			j = j.next
			k = k.next
		}
	}
	if j == nil {
		for i != nil {
			k.next = i
			i = i.next
			k = k.next
		}
	}
	return l
}

func (l *List) insert(v int) {
	node := &Node{val: v}
	if l.head == nil {
		l.head = node
		return
	}
	last := l.head
	for last.next != nil {
		last = last.next
	}
	last.next = node
}
func (l *List) insertAfter(find int, val int) {
	temp := l.head
	for temp.val != find {
		temp = temp.next
	}
	node := &Node{val: val, next: temp.next}
	temp.next = node
}
func (l *List) delete(val int) {

	if l.head.val == val {
		l.head = l.head.next
		return
	}

	find := l.head
	prev := find

	for find.val != val {
		prev = find
		find = find.next
	}
	prev.next = find.next
}
func (l *List) deleteNode(node *Node) {
	prev := node
	prev.val = node.next.val
	node = node.next
	prev.next = node.next

}
func (l *List) printlist() {
	last := l.head
	fmt.Print("list: ")
	for last != nil {
		fmt.Printf("%d ", last.val)
		last = last.next
	}
	fmt.Println("")
}
func (l *List) deletelist() {

	for l.head.next != nil {
		node := l.head.next
		l.head.next = node.next
		node.next = nil
	}
	l.head = nil
}
func (l *List) length() int {
	length := 0
	temp := l.head
	for temp != nil {
		length += 1
		temp = temp.next
	}
	fmt.Println("length: ", length)
	return length
}
func (l *List) find(val int) *Node {
	index := 0
	temp := l.head
	for temp.val != val {
		index += 1
		temp = temp.next
	}
	fmt.Printf("fonund at index: %d\n", index)
	return temp
}
func (l *List) reverse() {
	if l.length() == 1 {
		return
	}
	p, c, n := l.head, l.head.next, l.head.next.next
	for {
		c.next = p
		p = c
		c = n
		if n == nil {
			break
		}
		n = n.next
	}
	l.head.next = nil
	l.head = p
}
func (l *List) findmid() {
	s, f := l.head, l.head
	for (f.next != nil) && (f.next.next != nil) {
		s = s.next
		f = f.next.next
	}
	fmt.Println("mid = ", s.val)
}

func addNumber(a *List, b *List) *List {
	ans := new(List)

	i, j, k := a.head, b.head, ans.head
	for i != nil && j != nil {
		add := i.val + j.val
		if add/10 != 0 {
			if i.next == nil {
				node := new(Node)
				i.next = node
			}
			i.next.val += add / 10
		}
		node := &Node{val: add % 10}
		if ans.head == nil {
			ans.head = node
			k = ans.head
		}
		k.next = node
		k = k.next
		i = i.next
		j = j.next
	}
	if i == nil {
		for j != nil {
			node := &Node{val: j.val}
			k.next = node
			j = j.next
			k = k.next
		}
	}
	if j == nil {
		for i != nil {
			node := &Node{val: i.val}
			k.next = node
			i = i.next
			k = k.next
		}
	}

	return ans
}

func arrayTOList(a string) *List {
	l := new(List)

	aa := strings.Split(a, "")
	for _, v := range aa {
		vint, _ := strconv.Atoi(v)
		l.insert(vint)
	}

	return l
}

func findYjunction(a, b *List) bool {

	i := a.head
	for i != nil {
		j := b.head
		for j != nil {
			if j == i {
				return true
			}
			j = j.next
		}
		i = i.next
	}

	return false
}

func findIfLoop(a *List) bool {

	s, f := a.head, a.head
	for s != nil {
		if s == f {
			return true
		}
		s = s.next
		f = f.next.next
	}
	return false
}
func main() {

	a, b := arrayTOList("234"), arrayTOList("239")
	a.printlist()
	b.printlist()
	ans := addNumber(a, b)
	ans.printlist()

}
