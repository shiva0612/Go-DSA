package main

// creating new list
func mergeSortedLists(p, q *List) *List {
	var h, t *List

	for p != nil && q != nil {
		var y *List
		if p.Val <= q.Val {
			y = p
			p = p.Nxt
		} else {
			y = q
			q = q.Nxt
		}
		if h == nil {
			t = NewList(y.Val)
			h = t
		} else {
			t.Nxt = NewList(y.Val)
			t = t.Nxt
		}
	}
	var x *List
	if p == nil {
		x = q
	} else {
		x = p
	}
	for x != nil {
		t.Nxt = NewList(x.Val)
		t = t.Nxt
		x = x.Nxt
	}
	return h
}

// inplace
func mergeSortedLists2(p, q *List) *List {
	var x, y *List
	qt := q
	for qt != nil && p != nil {
		for qt.Val < p.Val {
			y = qt
			qt = qt.Nxt
		}
		x = p
		p = p.Nxt
		if y == nil {
			y = x
			q = x
			y.Nxt = qt
		} else {
			y.Nxt = x
			x.Nxt = qt
			y = x
		}
	}
	if p != nil {
		y.Nxt = p
	}
	return q
}
