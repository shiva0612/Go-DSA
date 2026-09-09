package main

func mathAdd(p, q *List) *List {

	var head *List
	t := head
	carry := 0
	for p != nil || q != nil {
		sum := carry
		if p != nil {
			sum += p.Val
			p = p.Nxt
		}

		if q != nil {
			sum += q.Val
			q = q.Nxt
		}

		var node *List
		if sum >= 10 {
			node = NewList(sum % 10)
			carry = sum / 10
		} else {
			node = NewList(sum)
			carry = 0
		}

		if head == nil {
			head = node
			t = head
		} else {
			t.Nxt = node
			t = t.Nxt
		}
	}
	if carry > 0 {
		t.Nxt = NewList(carry)
	}
	return head
}
