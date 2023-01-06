package linkedlist

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
func (l *List) findmid() *Node {
	s, f := l.head, l.head
	for (f.next != nil) && (f.next.next != nil) {
		s = s.next
		f = f.next.next
	}
	fmt.Println("mid = ", s.val)
	return s
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

func stringTOLList(a string) *List {
	l := new(List)

	aa := strings.Split(a, "")
	for _, v := range aa {
		vint, _ := strconv.Atoi(v)
		l.insert(vint)
	}

	return l
}

func (l *List) copy() *List {
	b := new(List)

	ltemp := l.head
	for ltemp != nil {
		b.insert(ltemp.val)
		ltemp = ltemp.next
	}

	return b
}
