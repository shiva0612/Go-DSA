package main

func reverseList(head *List) *List {
	c := head
	var p, n *List
	for c != nil {
		n = c.Nxt
		c.Nxt = p

		p = c
		c = n
	}
	return p
}
